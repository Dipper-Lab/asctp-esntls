package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transporter struct {
	Id            primitive.ObjectID `bson:"_id" validate:"required"`
	Name          string             `bson:"name" json:"name" validate:"required,min=5,max=15"`
	License       string             `bson:"license" json:"license"`
	Contact       string             `bson:"contact" json:"contact" validate:"required,min=10,max=10"`
	TransporterId string             `bson:"transporter_id"  json:"transporter_id"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}
