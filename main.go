package main

import (
	"context"
	"github.com/HardwareAndro/go-kanban-api/controller"
	"github.com/HardwareAndro/go-kanban-api/driver"
	"github.com/HardwareAndro/go-kanban-api/router"
	"github.com/HardwareAndro/go-kanban-api/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	r := gin.Default()

	projectDriver := driver.NewDriver()
	projectDriver.ConnectDatabase()
	defer func() {
		if err := projectDriver.Client.Disconnect(context.TODO()); err != nil {
			projectDriver.App.ErrorLogger.Println("Error disconnecting from MongoDB:", err)
		}
	}()
	projectService := service.NewProjectService(client)
	pc := controller.NewProjectController(projectService)
	routes := router.NewRouter(r, pc)

	routes.SetupRoutes()
	r.Run()
}
