package service

import (
	"context"
	"github.com/HardwareAndro/go-kanban-api/model"
	"github.com/HardwareAndro/go-kanban-api/repository"
	"log"
)

type CategoryService struct {
	categoryRepository repository.CategoryRepositoryI
	ctx                context.Context
}

func NewCategoryService(categoryRepository repository.CategoryRepositoryI, ctx context.Context) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
		ctx:                ctx,
	}
}

func (cs *CategoryService) GetCategories() []model.Category {
	category1 := model.Category{ID: "1"}
	category2 := model.Category{ID: "2"}
	category3 := model.Category{ID: "3"}

	return []model.Category{category1, category2, category3}
}

func (cs *CategoryService) AddCategory(ctx context.Context, category *model.Category) error {
	if ctx == nil {
		ctx = context.Background()
	}
	err := cs.categoryRepository.InsertData(context.Background(), category)
	if err != nil {
		return err
	}
	log.Println("Successfully Inserted Data User")

	return nil
}
