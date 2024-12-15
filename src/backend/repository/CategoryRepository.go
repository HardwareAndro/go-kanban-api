package repository

import (
	"context"
	"fmt"
	"github.com/HardwareAndro/go-kanban-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(collection *mongo.Collection) *CategoryRepository {
	return &CategoryRepository{
		collection: collection,
	}
}
func (cr *CategoryRepository) AddCategory(category *model.Category) (*model.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := cr.collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}
	insertedId := result.InsertedID
	if oid, ok := insertedId.(primitive.ObjectID); ok {
		category.ID = oid.Hex()
	} else {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}
	return category, nil
}
func (cr *CategoryRepository) GetCategories() ([]model.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := cr.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []model.Category

	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}
func (cr *CategoryRepository) GetCategoriesById(id string) (*model.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var category model.Category
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	err = cr.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}
func (cr *CategoryRepository) GetCategoryTasksById(id string) ([]model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var result struct {
		Tasks []model.Task `bson:"tasks"`
	}
	projection := bson.M{"tasks": 1, "_id": 0}
	err = cr.collection.FindOne(ctx, bson.M{"_id": objectId}, options.
		FindOne().
		SetProjection(projection)).
		Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Tasks, nil
}
func (cr *CategoryRepository) UpdateCategoryById(category *model.Category, id string) (*model.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{
			"name":  category.Name,
			"tasks": category.Tasks,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedCategory model.Category
	err = cr.collection.FindOneAndUpdate(ctx, bson.M{"_id": objectId}, update, opts).Decode(&updatedCategory)
	if err != nil {
		return nil, err
	}
	return &updatedCategory, nil
}
func (cr *CategoryRepository) DeleteCategoryById(id string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	deleteResult, err := cr.collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return 0, err
	}
	return deleteResult.DeletedCount, nil
}
