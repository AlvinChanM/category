package service

import (
	"category/domain/model"
	"category/domain/repository"
)


type ICategoryDataService interface {
	AddCategory( *model.Category)(int64, error)
	DeleteCategory(int64)error
	UpdateCategory(*model.Category)error
	FindCategoryByID(int64)(*model.Category, error)
	FindAllCategory()([]model.Category,error)
	FindCategoryByName(string)(*model.Category, error)
	FindCategoryByLevel(uint32)([]model.Category, error)
	FindCategoryByParent(int64)([]model.Category, error)
}

func NewCategoryDataService(repo repository.ICategoryRepository)ICategoryDataService{
	return &CategoryDataService{repo}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

func (c *CategoryDataService) AddCategory(category *model.Category) (int64, error) {
	return c.CategoryRepository.CreateCategory(category)
}

func (c *CategoryDataService) DeleteCategory(cid int64) error {
	return c.CategoryRepository.DeleteCategoryByID(cid)
}

func (c *CategoryDataService) UpdateCategory(category *model.Category) error {
	return c.CategoryRepository.UpdateCategory(category)
}

func (c *CategoryDataService) FindCategoryByID(cid int64) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByID(cid)
}

func (c *CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return c.CategoryRepository.FindAll()
}

func (c *CategoryDataService) FindCategoryByName(name string) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByName(name)
}

func (c *CategoryDataService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	return c.CategoryRepository.FindCategoryByLevel(level)
}

func (c *CategoryDataService) FindCategoryByParent(parentID int64) ([]model.Category, error) {
	return c.CategoryRepository.FindCategoryByParent(parentID)
}



