package controller

import (
	"github.com/HardwareAndro/go-kanban-api/service"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	service *service.ProjectService
}

func NewProjectController(ps *service.ProjectService) *ProjectController {
	return &ProjectController{
		service: ps,
	}
}

func (pc *ProjectController) GetProjects(ctx *gin.Context) {
	projects := pc.service.GetProjects()
	ctx.JSON(200, projects)
}
