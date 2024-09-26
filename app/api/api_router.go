package api

import (
	"github.com/HardwareAndro/go-kanban-api/app/api/controllers"
	"github.com/HardwareAndro/go-kanban-api/app/api/routers"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	router *gin.Engine
	pc     *controllers.ProjectController
	cc     *controllers.CategoryController
	tc     *controllers.TaskController
	uc     *controllers.UserController
}

func NewRouter(
	router *gin.Engine,
	pc *controllers.ProjectController,
	cc *controllers.CategoryController,
	tc *controllers.TaskController,
	uc *controllers.UserController,
) *ApiRouter {
	return &ApiRouter{
		router: router,
		pc:     pc,
		cc:     cc,
		tc:     tc,
		uc:     uc,
	}
}

func (ar *ApiRouter) SetupRoutes() {
	apiGroup := ar.router.Group("/api")

	routers.RegisterCategoryRoutes(apiGroup, ar.cc)
	routers.SetupProjectRoutes(apiGroup, ar.pc)
	routers.RegisterTaskRoutes(apiGroup, ar.tc)
	routers.SetupUserRoutes(ar.router, ar.uc)
}
