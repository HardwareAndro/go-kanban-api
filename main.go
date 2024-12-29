package main

import (
	"context"
	router "github.com/HardwareAndro/go-kanban-api/internal/api"
	"github.com/HardwareAndro/go-kanban-api/internal/api/controllers"
	"github.com/HardwareAndro/go-kanban-api/internal/api/services"
	model "github.com/HardwareAndro/go-kanban-api/internal/models"
	repository "github.com/HardwareAndro/go-kanban-api/internal/repositories"
	"github.com/HardwareAndro/go-kanban-api/internal/shared/constants"
	"github.com/HardwareAndro/go-kanban-api/pkg/mongo"
	ginzap "github.com/gin-contrib/zap"
	"go.uber.org/zap"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	projectDriver := mongo.NewDriver()
	projectDriver.ConnectDatabase()
	defer func() {
		if err := projectDriver.Client.Disconnect(context.TODO()); err != nil {
			zap.L().Error(constants.ERR_MONGO_CONNECTION, zap.String("Error", err.Error()))
		}
	}()
	categoryRepository := repository.NewGenericRepository[model.Category](projectDriver.CategoryColl)
	projectRepository := repository.NewGenericRepository[model.Project](projectDriver.ProjectColl)
	taskRepository := repository.NewGenericRepository[model.Task](projectDriver.TaskColl)
	userRepository := repository.NewGenericRepository[model.User](projectDriver.UserColl)

	categoryService := services.NewCategoryService(categoryRepository)
	projectService := services.NewProjectService(projectRepository)
	taskService := services.NewTaskService(taskRepository)
	userService := services.NewUserService(userRepository)

	cc := controllers.NewCategoryController(categoryService)
	pc := controllers.NewProjectController(projectService)
	tc := controllers.NewTaskController(taskService)
	uc := controllers.NewUserController(userService)

	routes := router.NewRouter(r, pc, cc, tc, uc)

	routes.SetupRoutes()
	r.Run(":8080")
}
