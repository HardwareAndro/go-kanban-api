package service

import "github.com/HardwareAndro/go-kanban-api/model"

type ProjectService struct{}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (ps *ProjectService) GetProjects() []model.Project {
	project1 := model.Project{ID: "1"}
	project2 := model.Project{ID: "2"}
	project3 := model.Project{ID: "3"}

	return []model.Project{project1, project2, project3}
}
func (ps *ProjectService) AddProject(project model.Project) {
	// TODO: After the database connection is made, add the project that comes here to the model
	var projects []model.Project
	projects = append(projects, project)
}
