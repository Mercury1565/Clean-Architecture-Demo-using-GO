package repository

import (
	"Clean_Architecture/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepo struct {
	database   mongo.Database
	collection string
}

func NewTaskRepo(database mongo.Database, collection string) domain.TaskRepository {
	return &taskRepo{
		database:   database,
		collection: collection,
	}
}

func (taskRepo *taskRepo) Create(c context.Context, task *domain.Task) error {
	collection := taskRepo.database.Collection(taskRepo.collection)

	_, err := collection.InsertOne(c, task)
	return err
}

func (taskRepo *taskRepo) FetchByUserID(c context.Context, taskID string) ([]domain.Task, error) {
	collection := taskRepo.database.Collection(taskRepo.collection)

	var tasks []domain.Task

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return tasks, err
	}

	cursor, err := collection.Find(c, bson.M{"userId": objID})
	if err != nil {
		return tasks, err
	}

	err = cursor.All(c, &tasks)
	if tasks == nil {
		return []domain.Task{}, err
	}

	return tasks, err
}
