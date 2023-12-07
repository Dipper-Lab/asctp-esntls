package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dipper-lab/asctp-esntls/consts"
	"github.com/dipper-lab/asctp-esntls/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrgSignedDetails struct {
	Id    string
	Name  string
	OrgId string
	Email string
	jwt.StandardClaims
}

var orgSecretKey = os.Getenv("ORG_SECRET_KEY")

func GenerateAllOrgTokens(id string, name string, orgId string, email string) (signedToken string, signedRefreshToken string, err error) {

	// now := time.Now().Local()

	claims := &OrgSignedDetails{
		Id:    id,
		Name:  name,
		OrgId: orgId,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Unix(),
			ExpiresAt: time.Now().Local().Add(time.Minute).Unix(),
		},
	}
	refreshClaims := &OrgSignedDetails{StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
	},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(orgSecretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(orgSecretKey))

	if err != nil {
		log.Panicln(err)
		return "", "", err
	}
	return token, refreshToken, err
}

func UpdateAllOrgTokens(signedToken string, signedRefreshToken string, email string) {

	var organisationCollection *mongo.Collection = database.OpenCollection(database.Client, consts.SystemDbName, "organisation")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_Token", Value: signedRefreshToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateObj = append(updateObj, bson.E{Key: "updated_At", Value: updatedAt})

	upsert := true
	filter := bson.M{"email": email}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := organisationCollection.UpdateOne(
		ctx,
		filter,
		bson.D{{Key: "$set", Value: updateObj}},
		&opt)

	defer cancel()

	if err != nil {
		log.Panicln(err)
	}
}

func ValidateOrgToken(signedToken string) (claims *OrgSignedDetails, msg string, valid bool) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&OrgSignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(orgSecretKey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*OrgSignedDetails)
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
