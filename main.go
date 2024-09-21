package main

import (
	"context"
	"fmt"
	"github.com/HardwareAndro/go-kanban-api/config"
	"github.com/HardwareAndro/go-kanban-api/controller"
	"github.com/HardwareAndro/go-kanban-api/driver"
	"github.com/HardwareAndro/go-kanban-api/router"
	"github.com/HardwareAndro/go-kanban-api/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var app config.GoAppTools
var client *mongo.Client

func ConnectDatabase() error {
	InfoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)

	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger

	err := godotenv.Load()
	if err != nil {
		app.ErrorLogger.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return fmt.Errorf("MONGODB_URI not found in environment")
	}

	client, err = driver.Connection(uri)
	if err != nil {
		return err
	}
	app.InfoLogger.Println("MongoDB's Database Connection Successfully Realized")

	return nil
}

func main() {
	err := ConnectDatabase()
	if err != nil {
		app.ErrorLogger.Fatalln("Failed to connect to the database:", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			app.ErrorLogger.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	r := gin.Default()

	projectService := service.NewProjectService(client)
	pc := controller.NewProjectController(projectService)
	routes := router.NewRouter(r, pc)

	routes.SetupRoutes()
	r.Run()
}
