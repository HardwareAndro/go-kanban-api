package main

import (
	"context"
	"github.com/HardwareAndro/go-kanban-api/controller"
	"github.com/HardwareAndro/go-kanban-api/driver"
	"github.com/HardwareAndro/go-kanban-api/repository"
	"github.com/HardwareAndro/go-kanban-api/router"
	"github.com/HardwareAndro/go-kanban-api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	projectDriver := driver.NewDriver()
	projectDriver.ConnectDatabase()
	defer func() {
		if err := projectDriver.Client.Disconnect(context.TODO()); err != nil {
			projectDriver.App.ErrorLogger.Println("Error disconnecting from MongoDB:", err)
		}
	}()
	categoryRepository := repository.NewCategoryRepository(projectDriver.CategoryColl)
	projectRepository := repository.NewProjectRepository(projectDriver.ProjectColl)

	categoryService := service.NewCategoryService(categoryRepository)
	projectService := service.NewProjectService(projectRepository)

	cc := controller.NewCategoryController(categoryService)
	pc := controller.NewProjectController(projectService)

	routes := router.NewRouter(r, pc, cc)

	routes.SetupRoutes()
	r.Run()
}
