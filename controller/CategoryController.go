package controller

import (
	"github.com/HardwareAndro/go-kanban-api/model"
	"github.com/HardwareAndro/go-kanban-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(cs *service.CategoryService) *CategoryController {
	return &CategoryController{
		service: cs,
	}
}

func (cc *CategoryController) GetCategories(ctx *gin.Context) {
	categories := cc.service.GetCategories()
	ctx.JSON(200, categories)
}

func (cc *CategoryController) AddCategory(ctx *gin.Context) {
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := cc.service.AddCategory(ctx, &category)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Successfully Inserted Data Category")
}
