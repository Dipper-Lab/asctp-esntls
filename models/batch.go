package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Batch struct {
	Id              primitive.ObjectID `bson:"_id" validate:"required"`
	BatchNumber     string             `bson:"batch_number" json:"batch_number"`
	ProduceId       string             `bson:"produce_id" json:"produce_id"`
	FarmBagIds      []string           `bson:"farmBags" json:"farmBags"`
	SupplierIds     []string           `bson:"suppliers" json:"suppliers"`
	State           string             `bson:"state" json:"state"`
	BatchId         string             `bson:"batch_id" json:"batch_id"`
	MoistureHistory []Moisture         `bson:"moisture" json:"moisture"`
	NfcUids         []string           `bson:"nfc_uid" json:"nfc_uid"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}
