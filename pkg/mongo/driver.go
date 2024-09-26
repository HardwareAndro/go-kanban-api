package mongo

import (
	"context"
	"github.com/HardwareAndro/go-kanban-api/internal/shared/constants"
	"go.uber.org/zap"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Driver struct {
	Client       *mongo.Client
	ProjectColl  *mongo.Collection
	CategoryColl *mongo.Collection
	TaskColl     *mongo.Collection
	UserColl     *mongo.Collection
}

func NewDriver() *Driver {
	return &Driver{}
}

func (dr *Driver) ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		zap.L().Error(constants.ERR_ENV_FILE_NOT_FOUND)
		return
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		zap.L().Error(constants.ERR_MONGODB_URI_NOT_FOUND)
		return
	}

	dr.Client, err = connection(uri)
	if err != nil {
		zap.L().Error(constants.ERR_MONGO_CONNECTION, zap.String("error", err.Error()))
		return
	}

	zap.L().Info(constants.SUCCESS_MONGODB_CONNECTION_ESTABLISHED)

	dr.ProjectColl = dr.Client.Database("go-kanban-api").Collection("projects")
	dr.CategoryColl = dr.Client.Database("go-kanban-api").Collection("categories")
	dr.TaskColl = dr.Client.Database("go-kanban-api").Collection("tasks")
	dr.UserColl = dr.Client.Database("go-kanban-api").Collection("users")
}

func connection(URI string) (*mongo.Client, error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()

	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
