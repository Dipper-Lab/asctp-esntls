package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Purchase struct {
	Id            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"   json:"name"`
	ProduceId     string             `bson:"produce_id"  json:"produce_id"`
	SupplierId    string             `bson:"supplier_id" json:"supplier_id"`
	BuyerUsername string             `bson:"buyer_username"  json:"buyer_username"`
	FarmerBagIds  []string           `bson:"farmerBag_ids" json:"farmerBag_ids"`
	PurchaseId    string             `bson:"purchase_id" json:"purchase_id"`
	CreatedAt     time.Time          `bson:"created_at"  json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"  json:"updated_at"`
}

//
