package main

import (
	"context"

	router "github.com/HardwareAndro/go-kanban-api/app/api"
	controller "github.com/HardwareAndro/go-kanban-api/app/api/controllers"
	service "github.com/HardwareAndro/go-kanban-api/app/api/services"
	model "github.com/HardwareAndro/go-kanban-api/app/models"
	"github.com/HardwareAndro/go-kanban-api/app/shared/constants"
	repository "github.com/HardwareAndro/go-kanban-api/app/shared/repositories"
	"github.com/HardwareAndro/go-kanban-api/app/shared/driver"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	projectDriver := driver.NewDriver()
	projectDriver.ConnectDatabase()
	defer func() {
		if err := projectDriver.Client.Disconnect(context.TODO()); err != nil {
			projectDriver.App.ErrorLogger.Println(constants.ERR_MONGO_CONNECTION, err)
		}
	}()
	categoryRepository := repository.NewGenericRepository[model.Category](projectDriver.CategoryColl)
	projectRepository := repository.NewGenericRepository[model.Project](projectDriver.ProjectColl)
	taskRepository := repository.NewGenericRepository[model.Task](projectDriver.TaskColl)
	userRepository := repository.NewGenericRepository[model.User](projectDriver.UserColl)

	categoryService := service.NewCategoryService(categoryRepository)
	projectService := service.NewProjectService(projectRepository)
	taskService := service.NewTaskService(taskRepository)
	userService := service.NewUserService(userRepository)

	cc := controller.NewCategoryController(categoryService)
	pc := controller.NewProjectController(projectService)
	tc := controller.NewTaskController(taskService)
	uc := controller.NewUserController(userService)

	routes := router.NewRouter(r, pc, cc, tc, uc)

	routes.SetupRoutes()
	r.Run(":8081")
}
