package models

import "time"

type FarmBag struct {
	Id              string               `bson:"_id" validate:"required"`
	BagNumber       string               `bson:"bag_number" json:"bag_number"` //format : MMDDYYXXXX
	ProduceId       string               `bson:"produce_id" json:"produce_id"`
	Weight          float32              `bson:"weight" json:"weight"`
	Price           float32              `bson:"price" json:"price"`
	MoistureHistory map[time.Time]string `bson:"moisture_history" json:"moisture_history"`
	SupplierId      string               `bson:"supplier_id" json:"supplier_id"`
	Status          string               `bson:"status" json:"status"` //unused, purchased, batched
	BatchId         string               `bson:"batch_id" json:"batch_id"`
	FarmBagId       string               `bson:"farmBag_id" json:"farmBag_id"`
	CreatedAt       time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time            `bson:"updated_at" json:"updated_at"`
}
