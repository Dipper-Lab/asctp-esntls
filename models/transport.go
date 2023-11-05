package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transport struct {
	Id            primitive.ObjectID `bson:"_id" validate:"required"`
	TransporterID string             `bson:"transporter_id" json:"transporter_id"`
	FacilityID    string             `bson:"facility_id" json:"facility_id"`
	NfcUids       []string           `bson:"nfc_uid" json:"nfc_uid"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
	Status        string             `bson:"status" json:"status"` //pending, initiated, arrived
}
