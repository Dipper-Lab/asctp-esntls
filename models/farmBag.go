package models

import "time"

type FarmBag struct {
	Id              string     `bson:"_id" validate:"required"`
	BagNumber       string     `bson:"bag_number" json:"bag_number"` //format : MMDDYYXXXX
	PurchaseId      string     `bson:"purchase_id" json:"purchase_id"`
	SupplierId      string     `bson:"supplier_id" json:"supplier_id"`
	ProduceId       string     `bson:"produce_id" json:"produce_id"`
	Weight          float32    `bson:"weight" json:"weight"`
	MoistureHistory []Moisture `bson:"moisture_history" json:"moisture_history"`
	Status          string     `bson:"status" json:"status"` //unused, purchased, batched
	BatchId         string     `bson:"batch_id" json:"batch_id"`
	FarmBagId       string     `bson:"farmBag_id" json:"farmBag_id"`
	CreatedAt       time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `bson:"updated_at" json:"updated_at"`
}
