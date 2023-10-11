package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sale struct {
	Id         primitive.ObjectID `bson:"_id"`
	SaleNumber string             `bson:"sale_number"  json:"sale_number"`
	CustomerId string             `bson:"customer_id"  json:"customer_id"`
	BatchId    string             `bson:"batch_id"   json:"batch_id"`
	CreatedAt  time.Time          `bson:"created_at"  json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"  json:"updated_at"`
}
