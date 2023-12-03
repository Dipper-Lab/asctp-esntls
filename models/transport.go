package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transport struct {
	Id            primitive.ObjectID `bson:"_id" validate:"required"`
	Name          string             `bson:"name" json:"name"`
	SourceID      string             `bson:"source_id" json:"source_id"`
	TransportID   string             `bson:"transport_id" json:"transport_id"`
	TransporterID string             `bson:"transporter_id" json:"transporter_id"`
	FacilityID    string             `bson:"facility_id" json:"facility_id"`
	NfcUids       []string           `bson:"nfc_uid" json:"nfc_uid"`
	InitiatedAt   time.Time          `bson:"initiated_at" json:"initiated_at"`
	ReceivedAt    time.Time          `bson:"received_at"  json:"received_at"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
	Status        string             `bson:"status" json:"status"` //pending, initiated, arrived
}
