package service

import (
	"github.com/HardwareAndro/go-kanban-api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectService struct{}

func NewProjectService(*mongo.Client) *ProjectService {
	return &ProjectService{}
}

func (ps *ProjectService) GetProjects() []model.Project {
	project1 := model.Project{ID: "1"}
	project2 := model.Project{ID: "2"}
	project3 := model.Project{ID: "3"}

	return []model.Project{project1, project2, project3}
}
