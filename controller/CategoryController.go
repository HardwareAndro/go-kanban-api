package controller

import (
	"errors"
	"github.com/HardwareAndro/go-kanban-api/model"
	"github.com/HardwareAndro/go-kanban-api/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
	categories, err := cc.service.GetCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get categories", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (cc *CategoryController) GetCategoryById(ctx *gin.Context) {
	id := ctx.Param("id")
	category, err := cc.service.GetCategoryById(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get category", "details": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (cc *CategoryController) GetCategoryTasksById(ctx *gin.Context) {
	id := ctx.Param("id")
	tasks, err := cc.service.GetCategoryTasksById(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Category's tasks not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get category's tasks", "details": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (cc *CategoryController) AddCategory(ctx *gin.Context) {
	var category model.Category
	if err := ctx.BindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := cc.service.AddCategory(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add category"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Category added successfully", "category": result})
}

func (cc *CategoryController) UpdateCategoryById(ctx *gin.Context) {
	id := ctx.Param("id")
	var category model.Category
	if err := ctx.BindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := cc.service.UpdateCategoryById(&category, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category by id"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Category update successfully", "category": result})
}

func (cc *CategoryController) DeleteCategoryById(ctx *gin.Context) {
	id := ctx.Param("id")
	deleteResult, err := cc.service.DeleteCategoryById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category", "details": err.Error()})
		return
	}
	if deleteResult == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
