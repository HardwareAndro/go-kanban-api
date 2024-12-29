package routers

import (
	"github.com/HardwareAndro/go-kanban-api/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(group *gin.RouterGroup, tc *controllers.TaskController) {
	tasksRouter := group.Group("/tasks")
	{
		tasksRouter.POST("/", tc.AddTask)
		tasksRouter.GET("/", tc.GetTasks)
		tasksRouter.GET("/:id", tc.GetTaskById)
		tasksRouter.PUT("/:id", tc.UpdateTaskById)
		tasksRouter.DELETE("/:id", tc.DeleteTaskById)
	}
}
