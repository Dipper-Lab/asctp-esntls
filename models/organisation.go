package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organisation struct {
	Id           primitive.ObjectID `bson:"_id" json:"id" validate:"required"`
	Name         string             `bson:"name"  json:"name" validate:"required,min=3,max=50"`
	Email        string             `bson:"email" json:"email"`
	OrgID        string             `bson:"org_id" json:"org_id" validate:"required,min=3,max=6"`
	Password     string             `bson:"password" json:"password" validate:"required"`
	Phone        string             `bson:"phone" json:"phone" validate:"required,min=10,max=10"`
	Token        *string            `bson:"token" json:"token"`
	RefreshToken *string            `bson:"refresh_Token" json:"refresh_Token"`
	CreatedAt    time.Time          `bson:"created_At" json:"created_At" validate:"required"`
	UpdatedAt    time.Time          `bson:"updated_At" json:"updated_At" validate:"required"`
}
