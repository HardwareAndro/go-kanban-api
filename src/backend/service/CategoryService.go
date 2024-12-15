package service

import (
	"github.com/HardwareAndro/go-kanban-api/config"
	"github.com/HardwareAndro/go-kanban-api/model"
	"github.com/HardwareAndro/go-kanban-api/repository"
	"log"
	"os"
)

type CategoryService struct {
	categoryRepository *repository.CategoryRepository
	App                config.GoAppTools
}

func NewCategoryService(categoryRepository *repository.CategoryRepository) *CategoryService {
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
func (cs *CategoryService) GetCategories() ([]model.Category, error) {
	categories, err := cs.categoryRepository.GetCategories()
	if err != nil {
		cs.App.ErrorLogger.Println("Failed to get categories", err)
		return nil, err
	}
	return categories, nil
}
func (cs *CategoryService) GetCategoryById(id string) (*model.Category, error) {
	category, err := cs.categoryRepository.GetCategoriesById(id)
	if err != nil {
		cs.App.ErrorLogger.Println("Failed to get categories by id", err)
		return nil, err
	}
	return category, nil
}
func (cs *CategoryService) GetCategoryTasksById(id string) ([]model.Task, error) {
	tasks, err := cs.categoryRepository.GetCategoryTasksById(id)
	if err != nil {
		cs.App.ErrorLogger.Println("Failed to get categoryTasks by id", err)
		return nil, err
	}
	return tasks, nil
}
func (cs *CategoryService) AddCategory(category *model.Category) (*model.Category, error) {
	result, err := cs.categoryRepository.AddCategory(category)
	if err != nil {
		cs.App.ErrorLogger.Println("Failed to add category", err)
		return nil, err
	}
	cs.App.InfoLogger.Println("Category added successfully with Id", result.ID)
	return result, nil
}
func (cs *CategoryService) UpdateCategoryById(category *model.Category, id string) (*model.Category, error) {
	result, err := cs.categoryRepository.UpdateCategoryById(category, id)
	if err != nil {
		cs.App.ErrorLogger.Println("Failed to update category by id", err)
		return nil, err
	}
	cs.App.InfoLogger.Println("Category updated successfully with id", result.ID)
	return result, nil
}
func (cs *CategoryService) DeleteCategoryById(id string) (int64, error) {
	result, err := cs.categoryRepository.DeleteCategoryById(id)
	if err != nil {
		cs.App.ErrorLogger.Println("Failed to delete category by id", err)
		return 0, err
	}
	cs.App.InfoLogger.Println("Category deleted successfully with Id: ", id)
	return result, err
}
