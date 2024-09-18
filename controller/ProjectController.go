package controller

import (
	"github.com/HardwareAndro/go-kanban-api/model"
	"github.com/HardwareAndro/go-kanban-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (pc *ProjectController) AddProject(ctx *gin.Context) {
	var project model.Project
	if err := ctx.BindJSON(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pc.service.AddProject(project)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Project added successfully", "project": project})
}
