package services

import (
	"github.com/HardwareAndro/go-kanban-api/internal/config"
	model2 "github.com/HardwareAndro/go-kanban-api/internal/models"
	repository "github.com/HardwareAndro/go-kanban-api/internal/repositories"
	"github.com/HardwareAndro/go-kanban-api/internal/shared/constants"
	"log"
	"os"
)

type CategoryService struct {
	categoryRepository *repository.GenericRepository[model2.Category]
	App                config.GoAppTools
}

func NewCategoryService(categoryRepository *repository.GenericRepository[model2.Category]) *CategoryService {
	InfoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)
	var app config.GoAppTools
	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger
	return &CategoryService{
		categoryRepository: categoryRepository,
		App:                app,
	}
}

func (cs *CategoryService) GetCategories() ([]model2.Category, error) {
	categories, err := cs.categoryRepository.FindAll()
	if err != nil {
		cs.App.ErrorLogger.Println(constants.ERR_CATEGORY_NOT_FOUND, err)
		return nil, err
	}
	return categories, nil
}

func (cs *CategoryService) GetCategoryById(id string) (*model2.Category, error) {
	category, err := cs.categoryRepository.FindById(id)
	if err != nil {
		cs.App.ErrorLogger.Println(constants.ERR_CATEGORY_NOT_FOUND, err)
		return nil, err
	}
	return category, nil
}

func (cs *CategoryService) GetCategoryTasksById(id string) ([]model2.Task, error) {
	// Since the GenericRepository doesn't support tasks directly, you will need a way to retrieve tasks.
	// If tasks are stored within a category, you can adapt your Category model to include a method for that.
	// Assuming that the `Tasks` field is available in the model.Category:
	category, err := cs.GetCategoryById(id)
	if err != nil {
		return nil, err
	}
	return category.Tasks, nil
}

func (cs *CategoryService) AddCategory(category *model2.Category) (*model2.Category, error) {
	_, err := cs.categoryRepository.Create(category)
	if err != nil {
		cs.App.ErrorLogger.Println(constants.ERR_ADD_CATEGORY, err)
		return nil, err
	}
	cs.App.InfoLogger.Println(constants.SUCCESS_ADD_CATEGORY, category.ID)
	return category, nil
}

func (cs *CategoryService) UpdateCategoryById(category *model2.Category, id string) (*model2.Category, error) {
	err := cs.categoryRepository.Update(id, category)
	if err != nil {
		cs.App.ErrorLogger.Println(constants.ERR_UPDATE_CATEGORY, err)
		return nil, err
	}
	cs.App.InfoLogger.Println(constants.SUCCESS_UPDATE_CATEGORY, id)
	return category, nil
}

func (cs *CategoryService) DeleteCategoryById(id string) (int64, error) {
	result, err := cs.categoryRepository.Delete(id)
	if err != nil {
		cs.App.ErrorLogger.Println(constants.SUCCESS_DELETE_CATEGORY, err)
		return 0, err
	}
	cs.App.InfoLogger.Println(constants.SUCCESS_DELETE_CATEGORY, id)
	return result, err
}
