package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Supplier struct {
	Id         primitive.ObjectID `bson:"_id" validate:"required"`
	Name       string             `bson:"name" json:"name" validate:"required,min=5,max=15"`
	Contact    string             `bson:"contact" json:"contact" validate:"required,min=10,max=10"`
	Email      string             `bson:"email" json:"email" `
	Location   string             `bson:"location" json:"location" validate:"required"`
	SupplierId string             `bson:"supplier_id" json:"supplier_id" validate:"required"`
	ProduceIds []string           `bson:"produce_ids" json:"produce_ids"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
