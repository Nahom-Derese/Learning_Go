package repositories

import (
	"context"
	"errors"

	"log"

	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/domain"
	"github.com/Nahom-Derese/Learning_Go/Task-8/task-manager/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection infrastructure.MongoCollection
}

// Delete implements TaskRepository.
func (repo *TaskRepository) Delete(c context.Context, id primitive.ObjectID) error {

	result, _ := repo.collection.DeleteOne(c, bson.M{"_id": id})

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

// FindAll implements TaskRepository.
func (repo *TaskRepository) FindAll(c context.Context, user domain.User) ([]domain.Task, error) {

	query := bson.D{}

	if user.Role != "admin" {
		query = bson.D{{Key: "user_id", Value: user.ID.Hex()}}
	}

	cursor, err := repo.collection.Find(c, query)

	if err != nil {
		log.Fatal(err)
	}

	var tasks []domain.Task = make([]domain.Task, 0)
	if err = cursor.All(c, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// FindByID implements TaskRepository.
func (repo *TaskRepository) FindByID(c context.Context, id primitive.ObjectID) (domain.Task, error) {

	var task domain.Task
	err := repo.collection.FindOne(c, bson.M{"_id": id}).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, errors.New("task not found")
		} else {
			log.Fatal(err)
		}
	}

	return task, nil
}

// Save implements TaskRepository.
func (repo *TaskRepository) Save(c context.Context, task domain.Task) (domain.Task, error) {

	InsertedTask, err := repo.collection.InsertOne(context.TODO(), task)

	if err != nil {
		return domain.Task{}, err
	}

	task.ID = InsertedTask.InsertedID.(primitive.ObjectID)
	return task, nil
}

// Constructor functions

func NewTaskRepository(collection infrastructure.MongoCollection) domain.TaskRepository {
	return &TaskRepository{collection}
}
