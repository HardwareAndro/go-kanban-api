package main

import (
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
	projectService := service.NewProjectService(client)
	pc := controller.NewProjectController(projectService)
	routes := router.NewRouter(r, pc)

	routes.SetupRoutes()
	r.Run()
}
