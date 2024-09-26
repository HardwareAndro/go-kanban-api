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

type ProjectRepository struct {
	collection *mongo.Collection
}

func NewProjectRepository(collection *mongo.Collection) *ProjectRepository {
	return &ProjectRepository{
		collection: collection,
	}
}
func (pr *ProjectRepository) CreateProject(project *model.Project) (*model.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := pr.collection.InsertOne(ctx, project)
	if err != nil {
		return nil, err
	}
	insertedId := result.InsertedID
	if oid, ok := insertedId.(primitive.ObjectID); ok {
		project.ID = oid.Hex()
	} else {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	return project, nil
}
func (pr *ProjectRepository) GetProjects() ([]model.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := pr.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var projects []model.Project
	if err = cursor.All(ctx, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}
func (pr *ProjectRepository) GetProjectById(id string) (*model.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var project model.Project
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = pr.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}
func (pr *ProjectRepository) GetProjectCategoriesById(id string) ([]model.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var result struct {
		Categories []model.Category `bson:"categories"`
	}
	projection := bson.M{"categories": 1, "_id": 0}
	err = pr.collection.FindOne(ctx, bson.M{"_id": objectId}, options.
		FindOne().
		SetProjection(projection)).
		Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Categories, nil
}
func (pr *ProjectRepository) UpdateProjectById(project *model.Project, id string) (*model.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{
			"name":       project.Name,
			"categories": project.Categories,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedProject model.Project
	err = pr.collection.FindOneAndUpdate(ctx, bson.M{"_id": objectId}, update, opts).Decode(&updatedProject)
	if err != nil {
		return nil, err
	}
	return &updatedProject, nil
}
func (pr *ProjectRepository) DeleteProjectById(id string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"_id": objectId}
	deleteResult, err := pr.collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return deleteResult.DeletedCount, nil
}
