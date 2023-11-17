package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	Id        primitive.ObjectID   `bson:"_id"`
	Name      string               `bson:"name"  json:"name"`
	Phone     string               `bson:"phone"  json:"phone"`
	Email     string               `bson:"email" json:"email" validate:"email"`
	Purchases map[string]time.Time `bson:"purchases"   json:"purchases"`
	/*
		{
			"batch_id":time
		}
	*/
	CustomerId string    `bson:"customer_id" json:"customer_id"`
	CreatedAt  time.Time `bson:"created_at"  json:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"  json:"updated_at"`
}
