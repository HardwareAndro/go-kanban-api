package service

import (
	"github.com/HardwareAndro/go-kanban-api/config"
	"github.com/HardwareAndro/go-kanban-api/model"
	"github.com/HardwareAndro/go-kanban-api/repository"
	"log"
	"os"
)

type ProjectService struct {
	projectRepository *repository.ProjectRepository
	App               config.GoAppTools
}

func NewProjectService(projectRepository *repository.ProjectRepository) *ProjectService {
	InfoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)
	var app config.GoAppTools
	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger
	return &ProjectService{
		projectRepository: projectRepository,
		App:               app,
	}
}

func (ps *ProjectService) GetProjects() ([]model.Project, error) {
	projects, err := ps.projectRepository.GetProjects()
	if err != nil {
		ps.App.ErrorLogger.Println("Failed to get projects", err)
		return nil, err
	}
	return projects, nil
}
func (ps *ProjectService) GetProjectById(id string) (*model.Project, error) {
	project, err := ps.projectRepository.GetProjectById(id)
	if err != nil {
		ps.App.ErrorLogger.Println("Failed to get projectById: ", err)
		return nil, err
	}
	return project, nil
}
func (ps *ProjectService) GetProjectCategoriesById(id string) ([]model.Category, error) {
	categories, err := ps.projectRepository.GetProjectCategoriesById(id)
	if err != nil {
		ps.App.ErrorLogger.Println("Failed to get projectCategoriesById: ", err)
		return nil, err
	}
	return categories, nil
}
func (ps *ProjectService) AddProject(project *model.Project) (*model.Project, error) {
	result, err := ps.projectRepository.CreateProject(project)
	if err != nil {
		ps.App.ErrorLogger.Fatalln("Failed to add Project", err)
		return nil, err
	}
	ps.App.InfoLogger.Println("Project added successfully with ID:", result.ID)
	return result, nil
}
func (ps *ProjectService) UpdateProjectById(project *model.Project, id string) (*model.Project, error) {
	updatedProject, err := ps.projectRepository.UpdateProjectById(project, id)
	if err != nil {
		ps.App.ErrorLogger.Println("Failed to update project: ", err)
		return nil, err
	}
	ps.App.InfoLogger.Println("Project updated successfully with Id: ", updatedProject.ID)
	return updatedProject, nil

}
func (ps *ProjectService) DeleteProjectById(id string) (int64, error) {
	deleteResult, err := ps.projectRepository.DeleteProjectById(id)
	if err != nil {
		ps.App.ErrorLogger.Println("Failed to delete projects: ", err)
		return 0, err
	}
	ps.App.InfoLogger.Println("Project deleted successfully with Id: ", id)
	return deleteResult, nil
}
