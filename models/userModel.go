package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName string             `json"first_name" validate:"required,min=3,max=30"`
	LastName  string             `json"last_name" validate:"required,min=3,max=30"`
	Age       int                `json:"age"`
	Gender    string             `json:"gender"omitempty`
	Email     string             `json:"email" validate:"email,required"`
	Phone     int                `json:"phone" validate:"email,required"`
}
