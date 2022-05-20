package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	userCollectionName = "user"
)

type Repository interface {
	UserRepo() UserRepo
}

type repositoryImpl struct {
	db *mongo.Database

	user UserRepo
}

func NewRepository(uri, dbName string) (Repository, error) {
	if client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri)); err == nil {
		r := &repositoryImpl{
			db: client.Database(dbName),
		}
		err = r.configure()
		if err != nil {
			return nil, err
		}
		r.init()
		return r, nil
	} else {
		return nil, fmt.Errorf("mongodb connect to %s failed: %s", uri, err.Error())
	}
}

func (r *repositoryImpl) UserRepo() UserRepo {
	return r.user
}

func (r *repositoryImpl) init() {
	r.user = NewUserRepo(r.db.Collection(userCollectionName))
}

func (r *repositoryImpl) configure() error {
	_, err := r.index(userCollectionName, true, "user_id")
	return err
}

func (r *repositoryImpl) index(collection string, unique bool, keys ...string) (string, error) {
	var composite bson.D
	for _, key := range keys {
		composite = append(composite, bson.E{
			Key:   key,
			Value: 1,
		})
	}
	return r.db.Collection(collection).Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    composite,
		Options: options.Index().SetUnique(unique),
	})
}
