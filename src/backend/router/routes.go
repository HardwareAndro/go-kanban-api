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

func NewRouter(router *gin.Engine, pc *controller.ProjectController, cc *controller.CategoryController) *Router {
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
		projectRouter.POST("/", r.pc.AddProject)
		projectRouter.GET("/:id", r.pc.GetProjectById)
		projectRouter.GET("/:id/categories", r.pc.GetProjectCategoriesById)
		projectRouter.PUT("/:id", r.pc.UpdateProjectById)
		projectRouter.DELETE("/:id", r.pc.DeleteProjectById)
	}

	categoryRouter := r.router.Group("/categories") // localhost:8080/categories
	{
		categoryRouter.POST("/", r.cc.AddCategory)
		categoryRouter.GET("/", r.cc.GetCategories)
		categoryRouter.GET("/:id", r.cc.GetCategoryById)
		categoryRouter.GET("/:id/tasks", r.cc.GetCategoryTasksById)
		categoryRouter.PUT("/:id", r.cc.UpdateCategoryById)
		categoryRouter.DELETE("/:id", r.cc.DeleteCategoryById)
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
