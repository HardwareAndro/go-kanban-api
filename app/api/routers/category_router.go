package routers

import (
	"github.com/HardwareAndro/go-kanban-api/app/api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(group *gin.RouterGroup, cc *controllers.CategoryController) {
	group.GET("/categories", cc.GetCategories)
	group.GET("/categories/:id", cc.GetCategoryById)
	group.POST("/categories", cc.AddCategory)
	group.PUT("/categories/:id", cc.UpdateCategoryById)
	group.DELETE("/categories/:id", cc.DeleteCategoryById)
}
