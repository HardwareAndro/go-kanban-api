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
	App    config.GoAppTools
	Client *mongo.Client
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
		dr.App.ErrorLogger.Println("No .env file found")
		return
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		dr.App.ErrorLogger.Fatalln(fmt.Errorf("MONGODB_URI not found in environment"))
		return
	}

	dr.Client, err = connection(uri)
	if err != nil {
		dr.App.ErrorLogger.Fatalln("Failed to connect to the database:", err)
		return
	}

	dr.App.InfoLogger.Println("MongoDB's Database Connection Successfully Realized")
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
