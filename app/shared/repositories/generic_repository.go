package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/HardwareAndro/go-kanban-api/app/shared/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GenericRepository[T any] struct {
	collection *mongo.Collection
}

func NewGenericRepository[T any](collection *mongo.Collection) *GenericRepository[T] {
	return &GenericRepository[T]{collection: collection}
}

// Create inserts a new document
func (gr *GenericRepository[T]) Create(entity *T) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(constants.TIMEOUT_DURATION)*time.Second)
	defer cancel()

	result, err := gr.collection.InsertOne(ctx, entity)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf(constants.ERR_INSERT_FAILED+": %w", err) // Use constant
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}
	return primitive.ObjectID{}, fmt.Errorf(constants.ERR_INVALID_OBJECT_ID) // Use constant
}

// FindAll retrieves all documents
func (gr *GenericRepository[T]) FindAll() ([]T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(constants.TIMEOUT_DURATION)*time.Second)
	defer cancel()

	cursor, err := gr.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var entities []T
	if err = cursor.All(ctx, &entities); err != nil {
		return nil, err
	}
	return entities, nil
}

// FindById retrieves a document by ID
func (gr *GenericRepository[T]) FindById(id string) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(constants.TIMEOUT_DURATION)*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf(constants.ERR_INVALID_OBJECT_ID+": %w", err)
	}

	var entity T
	err = gr.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&entity)
	if err != nil {
		return nil, fmt.Errorf(constants.ERR_ENTITY_NOT_FOUND+": %w", err)
	}
	return &entity, nil
}

func (gr *GenericRepository[T]) FindAsync(predicate func(bson.M) bson.M) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(constants.TIMEOUT_DURATION)*time.Second)
	defer cancel()

	var entity T
	err := gr.collection.FindOne(ctx, predicate(bson.M{})).Decode(&entity)
	if err != nil {
		return nil, fmt.Errorf(constants.ERR_ENTITY_NOT_FOUND+": %w", err)
	}
	return &entity, nil
}

// Update updates a document by ID
func (gr *GenericRepository[T]) Update(id string, entity *T) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(constants.TIMEOUT_DURATION)*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf(constants.ERR_INVALID_OBJECT_ID+": %w", err)
	}

	update := bson.M{"$set": entity}
	_, err = gr.collection.UpdateOne(ctx, bson.M{"_id": objectId}, update)
	if err != nil {
		return fmt.Errorf(constants.ERR_UPDATE_FAILED+": %w", err) // Use constant
	}
	return nil
}

// Delete removes a document by ID
func (gr *GenericRepository[T]) Delete(id string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(constants.TIMEOUT_DURATION)*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf(constants.ERR_INVALID_OBJECT_ID+": %w", err)
	}

	result, err := gr.collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return 0, fmt.Errorf(constants.ERR_DELETE_FAILED+": %w", err)
	}
	return result.DeletedCount, nil
}
