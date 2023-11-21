package models

import "time"

type Defect struct {
	Odour         bool      `bson:"odour" json:"odour" validate:"required"`
	Weevils       bool      `bson:"weevils" json:"weevils" validate:"required"`
	Contamination bool      `bson:"contamination" json:"contamination" validate:"required"`
	Tear          bool      `bson:"tear" json:"tear" validate:"required"`
	Notes         string    `bson:"notes" json:"notes" validate:"required,min=5"`
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updated_at"`
}
