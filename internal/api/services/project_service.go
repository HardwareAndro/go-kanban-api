package services

import (
	"github.com/HardwareAndro/go-kanban-api/internal/config"
	model2 "github.com/HardwareAndro/go-kanban-api/internal/models"
	repository "github.com/HardwareAndro/go-kanban-api/internal/repositories"
	"github.com/HardwareAndro/go-kanban-api/internal/shared/constants"
	"log"
	"os"
)

type ProjectService struct {
	projectRepository *repository.GenericRepository[model2.Project]
	App               config.GoAppTools
}

func NewProjectService(projectRepository *repository.GenericRepository[model2.Project]) *ProjectService {
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

func (ps *ProjectService) GetProjects() ([]model2.Project, error) {
	projects, err := ps.projectRepository.FindAll()
	if err != nil {
		ps.App.ErrorLogger.Println(constants.ERR_PROJECT_NOT_FOUND, err)
		return []model2.Project{}, nil // Return an empty slice instead of nil
	}
	return projects, nil
}

func (ps *ProjectService) GetProjectById(id string) (*model2.Project, error) {
	project, err := ps.projectRepository.FindById(id)
	if err != nil {
		ps.App.ErrorLogger.Println(constants.ERR_PROJECT_NOT_FOUND, err)
		return &model2.Project{}, nil
	}
	return project, nil
}

func (ps *ProjectService) GetProjectCategoriesById(id string) ([]model2.Category, error) {
	// Retrieve the project first, then access its categories
	project, err := ps.GetProjectById(id)
	if err != nil {
		return []model2.Category{}, nil
	}
	return project.Categories, nil // Assuming categories are directly part of the project model
}

func (ps *ProjectService) AddProject(project *model2.Project) (*model2.Project, error) {
	_, err := ps.projectRepository.Create(project)
	if err != nil {
		ps.App.ErrorLogger.Fatalln(constants.ERR_ADD_PROJECT, err)
		return nil, err
	}
	ps.App.InfoLogger.Println(constants.SUCCESS_ADD_PROJECT, project.ID)
	return project, nil
}

func (ps *ProjectService) UpdateProjectById(project *model2.Project, id string) (*model2.Project, error) {
	err := ps.projectRepository.Update(id, project)
	if err != nil {
		ps.App.ErrorLogger.Println(constants.ERR_UPDATE_PROJECT, err)
		return nil, err
	}
	ps.App.InfoLogger.Println(constants.SUCCESS_UPDATE_PROJECT, id)
	return project, nil
}

func (ps *ProjectService) DeleteProjectById(id string) (int64, error) {
	deleteResult, err := ps.projectRepository.Delete(id)
	if err != nil {
		ps.App.ErrorLogger.Println(constants.ERR_DELETE_PROJECT, err)
		return 0, err
	}
	ps.App.InfoLogger.Println(constants.SUCCESS_DELETE_PROJECT, id)
	return deleteResult, nil
}
