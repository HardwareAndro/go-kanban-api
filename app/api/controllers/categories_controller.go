package controllers

import (
	"errors"
	"net/http"

	service "github.com/HardwareAndro/go-kanban-api/app/api/services"
	model "github.com/HardwareAndro/go-kanban-api/app/models"
	"github.com/HardwareAndro/go-kanban-api/app/shared/constants"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_FAILED_TO_GET_CATEGORIES, "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (cc *CategoryController) GetCategoryById(ctx *gin.Context) {
	id := ctx.Param("id")
	category, err := cc.service.GetCategoryById(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": constants.ERR_CATEGORY_NOT_FOUND})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_FAILED_TO_GET_CATEGORY, "details": err.Error()})
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
			ctx.JSON(http.StatusNotFound, gin.H{"error": constants.ERR_TASK_NOT_FOUND})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_FAILED_TO_GET_CATEGORY_TASKS, "details": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_ADD_CATEGORY})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": constants.SUCCESS_ADD_CATEGORY, "category": result})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_UPDATE_CATEGORY})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": constants.SUCCESS_UPDATE_CATEGORY, "category": result})
}

func (cc *CategoryController) DeleteCategoryById(ctx *gin.Context) {
	id := ctx.Param("id")
	deleteResult, err := cc.service.DeleteCategoryById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_DELETE_CATEGORY})
		return
	}
	if deleteResult == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": constants.ERR_CATEGORY_NOT_FOUND})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": constants.SUCCESS_DELETE_CATEGORY})
}
