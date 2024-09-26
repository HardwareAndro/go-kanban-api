package services

import (
	"log"
	"os"

	model "github.com/HardwareAndro/go-kanban-api/app/models"
	"github.com/HardwareAndro/go-kanban-api/app/shared/constants"
	repository "github.com/HardwareAndro/go-kanban-api/app/shared/repositories"
	"github.com/HardwareAndro/go-kanban-api/config"
)

type ProjectService struct {
	projectRepository *repository.GenericRepository[model.Project]
	App               config.GoAppTools
}

func NewProjectService(projectRepository *repository.GenericRepository[model.Project]) *ProjectService {
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
	projects, err := ps.projectRepository.FindAll()
	if err != nil {
		ps.App.ErrorLogger.Println(constants.ERR_PROJECT_NOT_FOUND, err)
		return []model.Project{}, nil // Return an empty slice instead of nil
	}
	return projects, nil
}

func (ps *ProjectService) GetProjectById(id string) (*model.Project, error) {
	project, err := ps.projectRepository.FindById(id)
	if err != nil {
		ps.App.ErrorLogger.Println(constants.ERR_PROJECT_NOT_FOUND, err)
		return &model.Project{}, nil
	}
	return project, nil
}

func (ps *ProjectService) GetProjectCategoriesById(id string) ([]model.Category, error) {
	// Retrieve the project first, then access its categories
	project, err := ps.GetProjectById(id)
	if err != nil {
		return []model.Category{}, nil
	}
	return project.Categories, nil // Assuming categories are directly part of the project model
}

func (ps *ProjectService) AddProject(project *model.Project) (*model.Project, error) {
	_, err := ps.projectRepository.Create(project)
	if err != nil {
		ps.App.ErrorLogger.Fatalln(constants.ERR_ADD_PROJECT, err)
		return nil, err
	}
	ps.App.InfoLogger.Println(constants.SUCCESS_ADD_PROJECT, project.ID)
	return project, nil
}

func (ps *ProjectService) UpdateProjectById(project *model.Project, id string) (*model.Project, error) {
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
