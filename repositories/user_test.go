package repositories_test

import (
	"context"
	"encoding/json"
	"sample/entities"
	"sample/repositories"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userCollection *mongo.Collection
)

func init() {
	if client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017")); err == nil {
		userCollection = client.Database("sample").Collection("user")
	} else {
		panic("connect to database failed: " + err.Error())
	}
}

func TestUserCreate(t *testing.T) {
	repo := repositories.NewUserRepo(userCollection)
	ctx := context.Background()

	// in case user already exist
	_ = repo.Delele(ctx, "test_user_123")

	t.Run("create new user failed with empty user id", func(t *testing.T) {
		err := repo.Create(ctx, &entities.User{})
		assert.NotNil(t, err)
	})

	t.Run("create new user successfully", func(t *testing.T) {
		err := repo.Create(ctx, &entities.User{
			UserId:    "test_user_123",
			FirstName: "first_name",
			LastName:  "last_name",
			Nickname:  "nickname",
			Password:  "password",
			Email:     "email",
			Country:   "country",
		})
		assert.Nil(t, err)
	})

	t.Run("create duplicated user", func(t *testing.T) {
		err := repo.Create(ctx, &entities.User{
			UserId: "test_user_123",
		})
		assert.NotNil(t, err)
	})
}

func TestUserRead(t *testing.T) {
	repo := repositories.NewUserRepo(userCollection)
	ctx := context.Background()

	t.Run("read none exist user", func(t *testing.T) {
		_, err := repo.Read(ctx, "test_user_456")
		assert.NotNil(t, err)
	})

	t.Run("read user successfully", func(t *testing.T) {
		user, err := repo.Read(ctx, "test_user_123")
		assert.Nil(t, err)
		jsonStr := `{"user_id":"test_user_123","first_name":"first_name","last_name":"last_name","nickname":"nickname","password":"password","email":"email","country":"country"}`
		bytes, _ := json.Marshal(user)
		assert.JSONEq(t, jsonStr, string(bytes))
		assert.NotNil(t, user.CreatedAt)
		assert.Nil(t, user.UpdatedAt)
		assert.Equal(t, false, user.Id.IsZero())
	})
}

func TestUserUpdate(t *testing.T) {
	repo := repositories.NewUserRepo(userCollection)
	ctx := context.Background()

	t.Run("update none exist user", func(t *testing.T) {
		err := repo.Update(ctx, &entities.User{
			UserId: "test_user_456",
		})
		assert.NotNil(t, err)
	})

	t.Run("update user successfully", func(t *testing.T) {
		err := repo.Update(ctx, &entities.User{
			UserId:    "test_user_123",
			FirstName: "new_first_name",
			LastName:  "new_last_name",
			Nickname:  "new_nickname",
			Password:  "new_password",
			Email:     "new_email",
			Country:   "new_country",
		})
		assert.Nil(t, err)
		newUser, _ := repo.Read(ctx, "test_user_123")
		jsonStr := `{"user_id":"test_user_123","first_name":"new_first_name","last_name":"new_last_name","nickname":"new_nickname","password":"new_password","email":"new_email","country":"new_country"}`
		bytes, _ := json.Marshal(newUser)
		assert.JSONEq(t, jsonStr, string(bytes))
		assert.NotNil(t, newUser.CreatedAt)
		assert.NotNil(t, newUser.UpdatedAt)
		assert.Equal(t, false, newUser.Id.IsZero())
	})
}
