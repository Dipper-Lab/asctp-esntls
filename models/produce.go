package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Produce struct {
	Id                 primitive.ObjectID `bson:"_id"`
	Name               string             `bson:"name" json:"name"`
	Type               string             `bson:"type" json:"type"`
	QuantitiesMeasured map[string]string  `bson:"quantities_measured" json:"quantities_measured"` //{"quantity":"unit"}
	ProduceId          string             `bson:"produce_id" json:"produce_id"`
	CreatedAt          time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at" json:"updated_at"`
}
