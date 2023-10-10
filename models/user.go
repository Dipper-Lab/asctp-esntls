package models

import (
	"time"
)

type User struct {
	Id           *string   `bson:"_id" json:"id" validate:"required"`
	Username     string    `bson:"username" json:"username" validate:"required,min=5,max=50"`
	Password     string    `bson:"password" json:"password" validate:"required"`
	Role         string    `bson:"role" json:"role" validate:"required"`
	Facility     string    `bson:"facility" json:"facility" validate:"required"`
	FirstName    *string   `bson:"first_name" json:"first_name" validate:"required,min=2,max=50"`
	LastName     *string   `bson:"last_name" json:"last_name" validate:"required,min=2,max=50"`
	Email        *string   `bson:"email" json:"email" validate:"required,email"`
	Phone        string    `bson:"phone" json:"phone" validate:"required,min=10,max=10"`
	Token        *string   `bson:"token" json:"token"`
	RefreshToken *string   `bson:"refresh_Token" json:"refresh_Token"`
	CreatedAt    time.Time `bson:"created_At" json:"created_At" validate:"required"`
	UpdatedAt    time.Time `bson:"updated_At" json:"updated_At" validate:"required"`
}
