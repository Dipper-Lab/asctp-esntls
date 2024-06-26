package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Batch struct {
	Id              primitive.ObjectID `bson:"_id" validate:"required"`
	Name            string             `bson:"name" json:"name" validate:"required,min=2,max=15"`
	BatchNumber     string             `bson:"batch_number" json:"batch_number"`
	ProduceId       string             `bson:"produce_id" json:"produce_id"`
	FarmBagIds      []string           `bson:"farmBags" json:"farmBags"`
	SupplierIds     []string           `bson:"suppliers" json:"suppliers"`
	Status          string             `bson:"status" json:"status"`
	BatchId         string             `bson:"batch_id" json:"batch_id"`
	MoistureHistory []Moisture         `bson:"moisture" json:"moisture"`
	NfcUids         []string           `bson:"nfc_uid" json:"nfc_uid"`
	Notes           []Note             `bson:"notes" json:"notes"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}
