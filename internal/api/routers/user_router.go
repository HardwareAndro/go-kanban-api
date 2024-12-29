package routers

import (
	"github.com/HardwareAndro/go-kanban-api/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, uc *controllers.UserController) {
	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", uc.RegisterUser)
		userRouter.POST("/login", uc.LoginUser)
		userRouter.POST("/logout", uc.LogoutUser)
	}
}
