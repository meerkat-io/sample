package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `bson:"_id"`

	UserId    string `bson:"user_id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Nickname  string `bson:"nickname"`
	Password  string `bson:"password"`
	Email     string `bson:"email"`
	Country   string `bson:"country"`

	CreatedAt *time.Time `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"`
}
