package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dipper-lab/asctp-esntls/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserSignedDetails struct {
	Id       string
	Username string
	Role     string
	Facility string
	OrgId    string
	jwt.StandardClaims
}

var userSecretKey = os.Getenv("USER_SECRET_KEY")

func GenerateAllTokens(id string, username string, role string, facility string, orgId string) (signedToken string, signedRefreshToken string, err error) {

	now := time.Now().Local()

	claims := &UserSignedDetails{
		Id:       id,
		Username: username,
		Role:     role,
		Facility: facility,
		OrgId:    orgId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Unix(),
		},
	}
	refreshClaims := &UserSignedDetails{StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
	},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(userSecretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(userSecretKey))

	if err != nil {
		log.Panicln(err)
		return "", "", err
	}
	return token, refreshToken, err
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, dbName string, userId string) {

	var userCollection *mongo.Collection = database.OpenCollection(database.Client, dbName, "user")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_Token", Value: signedRefreshToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateObj = append(updateObj, bson.E{Key: "updated_At", Value: updatedAt})

	upsert := true
	filter := bson.M{"_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{{Key: "$set", Value: updateObj}},
		&opt)

	defer cancel()

	if err != nil {
		log.Panicln(err)
	}
}

func ValidateToken(signedToken string) (claims *UserSignedDetails, msg string, valid bool) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&UserSignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(userSecretKey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*UserSignedDetails)
	validity := true
	if !ok {
		_ = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		validity = false
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		_ = fmt.Sprintf("token is expired")
		msg = err.Error()
		validity = false
		return
	}
	return claims, msg, validity
}
