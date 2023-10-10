package helpers

import (
	"context"
	"time"

	"github.com/dipper-lab/asctp-esntls/database"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfExistsInCollection(dbName string, collectionName string, key string, value string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	collection := database.OpenCollection(database.Client, dbName, collectionName)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{key: value})
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}
