package controllers

import (
	"errors"
	service "github.com/HardwareAndro/go-kanban-api/internal/api/services"
	"github.com/HardwareAndro/go-kanban-api/internal/models"
	"github.com/HardwareAndro/go-kanban-api/internal/shared/constants"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_FAILED_TO_GET_PROJECT, "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, projects)
}

func (pc *ProjectController) GetProjectById(ctx *gin.Context) {
	id := ctx.Param("id")
	project, err := pc.service.GetProjectById(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": constants.ERR_PROJECT_NOT_FOUND})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_FAILED_TO_GET_PROJECT, "details": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, project)
}

func (pc *ProjectController) GetProjectCategoriesById(ctx *gin.Context) {
	id := ctx.Param("id")
	categories, err := pc.service.GetProjectCategoriesById(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": constants.ERR_CATEGORY_NOT_FOUND})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_FAILED_TO_GET_PROJECT_CATEGORIES, "details": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_ADD_PROJECT})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": constants.SUCCESS_ADD_PROJECT, "project": result})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_UPDATE_PROJECT})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": constants.SUCCESS_UPDATE_PROJECT, "project": updatedProject})
}

func (pc *ProjectController) DeleteProjectById(ctx *gin.Context) {
	id := ctx.Param("id")
	deleteResult, err := pc.service.DeleteProjectById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_DELETE_PROJECT, "details": err.Error()})
		return
	}
	if deleteResult == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": constants.ERR_PROJECT_NOT_FOUND})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": constants.SUCCESS_DELETE_PROJECT})
}
