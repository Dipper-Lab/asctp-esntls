package helpers

import (
	"context"
	"net/http"
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

func ValidateProperty(dbName string, collectionName string, propertyKey string, propertyValue string) (bool, int, string) {
	isPropertyValid, err := CheckIfExistsInCollection(dbName, collectionName, propertyKey, propertyValue)
	if err != nil {
		return false, http.StatusInternalServerError, "error occurred whilst validating " + propertyKey

	}

	if !isPropertyValid {
		return false, http.StatusBadRequest, propertyKey + " is invalid"
	}

	return true, 0, ""
}

func ValidatePropertyList(dbName string, collectionName string, propertyKey string, propertyValues []string) (bool, int, string) {

	for _, propertyValue := range propertyValues {
		isPropertyValid, err := CheckIfExistsInCollection(dbName, collectionName, propertyKey, propertyValue)
		if err != nil {
			return false, http.StatusInternalServerError, "error occurred whilst validating produce id"

		}
		if !isPropertyValid {
			return false, http.StatusBadRequest, propertyKey + " is invalid"
		}
	}

	return true, 0, ""
}

func ValidatePropertiesWithMap(dbName string, collectionName string, properties map[string]string) (bool, int, string) {

	for propertyKey, propertyValue := range properties {
		isPropertyValid, err := CheckIfExistsInCollection(dbName, collectionName, propertyKey, propertyValue)
		if err != nil {
			return false, http.StatusInternalServerError, "error occurred whilst validating produce id"

		}
		if !isPropertyValid {
			return false, http.StatusBadRequest, propertyKey + " is invalid"
		}
	}

	return true, 0, ""
}

//waste comment
