package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id" validate:"required"`
	Username     string             `bson:"username" json:"username" validate:"required,min=5,max=50"`
	Password     string             `bson:"password" json:"password" validate:"required"`
	Role         string             `bson:"role" json:"role" validate:"required"`
	OrgId        string             `bson:"org_id" json:"org_id"`
	Facility     string             `bson:"facility" json:"facility" validate:"required"`
	FirstName    *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=50"`
	LastName     *string            `bson:"last_name" json:"last_name" validate:"required,min=2,max=50"`
	Email        *string            `bson:"email" json:"email" validate:"required,email"`
	Phone        string             `bson:"phone" json:"phone" validate:"required,min=10,max=10"`
	Token        *string            `bson:"token" json:"token"`
	RefreshToken *string            `bson:"refresh_Token" json:"refresh_Token"`
	IsActive     bool               `bson:"isActive" json:"isActive"`
	CreatedAt    time.Time          `bson:"created_At" json:"created_At" validate:"required"`
	UpdatedAt    time.Time          `bson:"updated_At" json:"updated_At" validate:"required"`
}
