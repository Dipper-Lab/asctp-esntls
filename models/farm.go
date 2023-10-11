package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Farm struct {
	Id         primitive.ObjectID `bson:"_id" validate:"required"`
	Name       string             `bson:"name" json:"name" validate:"required,min=5,max=15"`
	Contact    string             `bson:"contact" json:"contact" validate:"required,min=10,max=10"`
	ProduceIds []string           `bson:"produce" json:"produce"`
	FarmId     string             `bson:"farm_id" json:"farm_id" validate:"required"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
