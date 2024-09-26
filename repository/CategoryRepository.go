package repository

import (
	"context"
	"github.com/HardwareAndro/go-kanban-api/model"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type CategoryRepository struct {
	mongoDB *mongo.Database
}

type CategoryRepositoryI interface {
	InsertData(ctx context.Context, req *model.Category) error
}

func NewCategoryRepository(mongoDB *mongo.Database) CategoryRepositoryI {
	return &CategoryRepository{
		mongoDB: mongoDB,
	}
}

func (c CategoryRepository) InsertData(ctx context.Context, req *model.Category) error {

	_, err := c.mongoDB.Collection("categories").InsertOne(ctx, req)
	if err != nil {
		log.Println("error")
	}

	return err
}
