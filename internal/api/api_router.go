package api

import (
	controllers2 "github.com/HardwareAndro/go-kanban-api/internal/api/controllers"
	routers2 "github.com/HardwareAndro/go-kanban-api/internal/api/routers"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	router *gin.Engine
	pc     *controllers2.ProjectController
	cc     *controllers2.CategoryController
	tc     *controllers2.TaskController
	uc     *controllers2.UserController
}

func NewRouter(
	router *gin.Engine,
	pc *controllers2.ProjectController,
	cc *controllers2.CategoryController,
	tc *controllers2.TaskController,
	uc *controllers2.UserController,
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

	routers2.RegisterCategoryRoutes(apiGroup, ar.cc)
	routers2.SetupProjectRoutes(apiGroup, ar.pc)
	routers2.RegisterTaskRoutes(apiGroup, ar.tc)
	routers2.SetupUserRoutes(ar.router, ar.uc)
}
