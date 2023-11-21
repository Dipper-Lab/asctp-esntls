package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Defect struct {
	Id            primitive.ObjectID `bson:"_id"`
	Odour         bool               `bson:"odour" json:"odour" validate:"required"`
	Weevils       bool               `bson:"weevils" json:"weevils" validate:"required"`
	Contamination bool               `bson:"contamination" json:"contamination" validate:"required"`
	Tear          bool               `bson:"tear" json:"tear" validate:"required"`
	Notes         string             `bson:"notes" json:"notes" validate:"required,min=5"`
	NfCUid        string             `bson:"nfc_uid" json:"nfc_uid" validate:"required"`
	BatchId       string             `bson:"batch_id" json:"batch_id" validate:"required"`
	DefectId      string             `bson:"defect_id" json:"defect_id" validate:"required"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}
