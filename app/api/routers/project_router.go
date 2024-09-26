package routers

import (
	"github.com/HardwareAndro/go-kanban-api/app/api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupProjectRoutes(group *gin.RouterGroup, pc *controllers.ProjectController) {
	group.GET("/projects", pc.GetProjects)
	group.GET("/projects/:id", pc.GetProjectById)
	group.POST("/projects", pc.AddProject)
	group.PUT("/projects/:id", pc.UpdateProjectById)
	group.DELETE("/projects/:id", pc.DeleteProjectById)
}
