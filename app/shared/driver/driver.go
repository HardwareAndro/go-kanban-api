package driver

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HardwareAndro/go-kanban-api/app/shared/constants"
	"github.com/HardwareAndro/go-kanban-api/config"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Driver struct {
	App          config.GoAppTools
	Client       *mongo.Client
	ProjectColl  *mongo.Collection
	CategoryColl *mongo.Collection
	TaskColl     *mongo.Collection
	UserColl     *mongo.Collection
}

func NewDriver() *Driver {
	InfoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)
	var app config.GoAppTools
	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger
	return &Driver{
		App: app,
	}
}

func (dr *Driver) ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		dr.App.ErrorLogger.Println(constants.ERR_ENV_FILE_NOT_FOUND)
		return
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		dr.App.ErrorLogger.Fatalln(fmt.Errorf(constants.ERR_MONGODB_URI_NOT_FOUND))
		return
	}

	dr.Client, err = connection(uri)
	if err != nil {
		dr.App.ErrorLogger.Fatalln(constants.ERR_MONGO_CONNECTION, err)
		return
	}
	dr.App.InfoLogger.Println(constants.SUCCESS_MONGODB_CONNECTION_ESTABLISHED)

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
