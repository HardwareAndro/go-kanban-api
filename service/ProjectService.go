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
func (ps *ProjectService) GetProjectCategoriesById(id string) []model.Category {
	// TODO: After the database connection is made, return the Categories with the id found in the project model
	projects := []model.Project{
		{ID: "1", Name: "Project 1", Categories: []model.Category{}},
		{ID: "2", Name: "Project 2", Categories: []model.Category{}},
		{ID: "3", Name: "Project 3", Categories: []model.Category{}},
	}
	for _, project := range projects {
		if project.ID == id {
			return project.Categories
		}
	}
	return []model.Category{}
}
