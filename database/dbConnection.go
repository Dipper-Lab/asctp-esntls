package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_ATLAS_URL")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return mongoClient
}

var Client *mongo.Client = DbInstance()

func OpenCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {

	err := godotenv.Load(filepath.Join(rootDir(), ".env"))
	// err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var collection *mongo.Collection = client.Database(dbName).Collection(collectionName)

	return collection
}

func CheckIfDbExistsInCluster(client *mongo.Client, dbName string) bool {

	dbs, err := client.ListDatabases(context.Background(), bson.M{})
	if err != nil {
		log.Panic(err)
	}

	for _, db := range dbs.Databases {
		if db.Name == dbName {
			return true
		}
	}

	return false

}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
