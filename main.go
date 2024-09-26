package main

import (
	"context"
	"fmt"
	"github.com/HardwareAndro/go-kanban-api/config"
	"github.com/HardwareAndro/go-kanban-api/controller"
	"github.com/HardwareAndro/go-kanban-api/repository"
	"github.com/HardwareAndro/go-kanban-api/router"
	"github.com/HardwareAndro/go-kanban-api/service"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	ctx context.Context
)

func main() {
	r := gin.Default()
	ctx = context.TODO()

	mongoConnection, err := config.ConnectDB()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("mongo connection established")

	cr := repository.NewCategoryRepository(mongoConnection)
	cs := service.NewCategoryService(cr, ctx)
	cc := controller.NewCategoryController(cs)

	ps := service.NewProjectService()
	pc := controller.NewProjectController(ps)

	routes := router.NewRouter(r, pc, cc)

	routes.SetupRoutes()

	r.Run()
}
