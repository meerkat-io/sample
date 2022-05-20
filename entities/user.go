package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `bson:"_id" json:"-"`

	UserId    string `bson:"user_id" json:"user_id"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Nickname  string `bson:"nickname" json:"nickname"`
	Password  string `bson:"password" json:"password"`
	Email     string `bson:"email" json:"email"`
	Country   string `bson:"country" json:"country"`

	CreatedAt *time.Time `bson:"created_at" json:"-"`
	UpdatedAt *time.Time `bson:"updated_at" json:"-"`
}
