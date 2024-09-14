package main

import (
	"github.com/HardwareAndro/go-kanban-api/controller"
	"github.com/HardwareAndro/go-kanban-api/router"
	"github.com/HardwareAndro/go-kanban-api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	projectService := service.NewProjectService()
	pc := controller.NewProjectController(projectService)
	routes := router.NewRouter(r, pc)

	routes.SetupRoutes()
	r.Run()
}
