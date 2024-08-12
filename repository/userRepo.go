package repository

import (
	"Clean_Architecture/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	database   mongo.Database
	collection string
}

func NewUserRepo(database mongo.Database, collection string) domain.UserRepository {
	return &userRepo{
		database:   database,
		collection: collection,
	}
}

func (userRepo *userRepo) Create(c context.Context, user *domain.User) error {
	collection := userRepo.database.Collection(userRepo.collection)

	_, err := collection.InsertOne(c, user)
	return err
}

func (userRepo *userRepo) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := userRepo.database.Collection(userRepo.collection)

	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (userRepo *userRepo) GetByID(c context.Context, userID string) (domain.User, error) {
	collection := userRepo.database.Collection(userRepo.collection)

	var user domain.User
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		fmt.Println("heloooo")
		return user, err
	}

	return user, err
}
