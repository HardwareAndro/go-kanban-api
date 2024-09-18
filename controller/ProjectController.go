package controller

import (
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
func (pc *ProjectController) GetProjectCategoriesById(ctx *gin.Context) {
	id := ctx.Param("id")
	categories := pc.service.GetProjectCategoriesById(id)
	ctx.JSON(http.StatusOK, categories)
}
