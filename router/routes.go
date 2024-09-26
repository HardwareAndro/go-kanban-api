package router

import (
	"github.com/HardwareAndro/go-kanban-api/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	pc     *controller.ProjectController
	cc     *controller.CategoryController
}

func NewRouter(router *gin.Engine,
	pc *controller.ProjectController,
	cc *controller.CategoryController) *Router {
	return &Router{
		router: router,
		pc:     pc,
		cc:     cc,
	}
}

func (r *Router) SetupRoutes() {
	projectRouter := r.router.Group("/projects") // localhost:8080/projects/1231/categories
	{
		projectRouter.GET("/", r.pc.GetProjects)
		projectRouter.POST("/", func(context *gin.Context) {})
		projectRouter.GET("/:id", func(context *gin.Context) {})
		projectRouter.GET("/:id/categories", func(context *gin.Context) {})
		projectRouter.PUT("/:id", func(context *gin.Context) {})
		projectRouter.DELETE("/:id", func(context *gin.Context) {})
	}

	categoryRouter := r.router.Group("/categories") // localhost:8080/categories
	{
		categoryRouter.GET("/", r.cc.GetCategories)
		categoryRouter.POST("/", r.cc.AddCategory)
		categoryRouter.GET("/:id", func(context *gin.Context) {})
		categoryRouter.GET("/:id/tasks", func(context *gin.Context) {})
		categoryRouter.PUT("/:id", func(context *gin.Context) {})
		categoryRouter.DELETE("/:id", func(context *gin.Context) {})
	}

	tasksRouter := r.router.Group("/tasks") // localhost:8080/tasks
	{
		tasksRouter.POST("/", func(context *gin.Context) {})
		tasksRouter.GET("/", func(context *gin.Context) {})
		tasksRouter.GET("/:id", func(context *gin.Context) {})
		tasksRouter.PUT("/:id", func(context *gin.Context) {})
		tasksRouter.DELETE("/:id", func(context *gin.Context) {})
	}

	userRouter := r.router.Group("/users") // localhost:8080/users
	{
		userRouter.POST("/register", func(context *gin.Context) {})
		userRouter.POST("/login", func(context *gin.Context) {})
		userRouter.POST("/logout", func(context *gin.Context) {})
	}
}
