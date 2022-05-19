package repositories

import (
	"context"
	"time"

	"sample/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	Create(ctx context.Context, user *entities.User) error
	Read(ctx context.Context, userId string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delele(ctx context.Context, userId string) error
}

type userRepoImpl struct {
	collection *mongo.Collection
}

func (u *userRepoImpl) Create(ctx context.Context, user *entities.User) error {
	if user.Id.IsZero() {
		user.Id = primitive.NewObjectID()
	}
	t := time.Now()
	user.CreatedAt = &t
	_, err := u.collection.InsertOne(ctx, user)
	return err
}

func (u *userRepoImpl) Read(ctx context.Context, userId string) (*entities.User, error) {
	user := &entities.User{}
	err := u.collection.FindOne(ctx, bson.M{
		"user_id": userId,
	}).Decode(user)
	return user, err
}

func (u *userRepoImpl) Update(ctx context.Context, user *entities.User) error {
	exist, err := u.Read(ctx, user.UserId)
	if err != nil {
		return err
	}
	user.Id = exist.Id
	user.CreatedAt = exist.CreatedAt
	t := time.Now()
	user.UpdatedAt = &t
	_, err = u.collection.ReplaceOne(context.Background(), bson.M{"_id": user.Id}, user)
	return err
}

func (u *userRepoImpl) Delele(ctx context.Context, userId string) error {
	result, err := u.collection.DeleteOne(ctx, bson.M{
		"user_id": userId,
	})
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return err
}
