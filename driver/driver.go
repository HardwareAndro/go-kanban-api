package driver

import (
	"context"
	"fmt"
	"github.com/HardwareAndro/go-kanban-api/config"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Driver struct {
	app    config.GoAppTools
	client *mongo.Client
}

func NewDriver() *Driver {
	InfoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)
	var app config.GoAppTools
	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger
	return &Driver{
		app: app,
	}
}
func (dr *Driver) ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		dr.app.ErrorLogger.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		dr.app.ErrorLogger.Fatalln(fmt.Errorf("MONGODB_URI not found in environment"))
	}

	dr.client, err = connection(uri)
	if err != nil {
		dr.app.ErrorLogger.Fatalln("Failed to connect to the database:", err)
	}
	defer func() {
		if err := dr.client.Disconnect(context.TODO()); err != nil {
			dr.app.ErrorLogger.Println("Error disconnecting from MongoDB:", err)
		}
	}()
	dr.app.InfoLogger.Println("MongoDB's Database Connection Successfully Realized")
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
