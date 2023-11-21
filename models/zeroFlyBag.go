package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ZeroFlyBag struct {
	Id        primitive.ObjectID              `bson:"_id"`
	BagNumber string                          `bson:"bag_number" json:"bag_number"`
	NFCUID    string                          `bson:"nfc_uid" json:"nfc_uid"`
	History   map[string]map[string]time.Time `bson:"history" json:"history"`
	/*
		batch_id:{
			attached_at: time.Time,
			detached_at:time.Time
		}
	*/
	Status    string    `bson:"status" json:"status"`
	Location  string    `bson:"location" json:"location"`
	CreatedAt time.Time `bson:"created_at"   json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
