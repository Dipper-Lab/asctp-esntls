package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SensorReadings struct {
	Id          primitive.ObjectID `bson:"_id"`
	Temperature float32            `bson:"temperature" json:"temperature"`
	Humidity    float32            `bson:"humidity" json:"humidity"`
	Location    string             `bson:"location"  json:"location"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
