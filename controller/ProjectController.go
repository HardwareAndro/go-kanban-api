package controller

import (
	"errors"
	"github.com/HardwareAndro/go-kanban-api/model"
	"github.com/HardwareAndro/go-kanban-api/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
	projects, err := pc.service.GetProjects()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get projects", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, projects)
}

func (pc *ProjectController) GetProjectById(ctx *gin.Context) {
	id := ctx.Param("id")

	project, err := pc.service.GetProjectById(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project", "details": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, project)
}

func (pc *ProjectController) GetProjectCategoriesById(ctx *gin.Context) {
	id := ctx.Param("id")

	categories, err := pc.service.GetProjectById(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Project Categories not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project categories", "details": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (pc *ProjectController) AddProject(ctx *gin.Context) {
	var project model.Project
	if err := ctx.BindJSON(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := pc.service.AddProject(&project)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add project"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Project added successfully", "project": result})
}

func (pc *ProjectController) UpdateProjectById(ctx *gin.Context) {
	id := ctx.Param("id")
	var project model.Project
	if err := ctx.BindJSON(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedProject, err := pc.service.UpdateProjectById(&project, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Project updated successfully", "project": updatedProject})
}

func (pc *ProjectController) DeleteProjectById(ctx *gin.Context) {
	id := ctx.Param("id")
	deleteResult, err := pc.service.DeleteProjectById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project", "details": err.Error()})
		return
	}
	if deleteResult == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
