package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Facility struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name" json:"name"`
	Type       string             `bson:"type" json:"type"`
	Location   string             `bson:"location" json:"location"`
	Capacity   string             `bson:"capacity" json:"capacity"`
	FacilityId string             `bson:"facility_id"  json:"facility_id"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
