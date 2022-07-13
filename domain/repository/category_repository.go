package repository

import (
	"category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(*model.Category)(int64, error)
	DeleteCategoryByID(int64)error
	UpdateCategory(*model.Category)error
	FindAll() ([]model.Category, error)
	FindCategoryByName(string2 string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}


func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDB: db}
}

type CategoryRepository struct {
	mysqlDB *gorm.DB
}

func (c *CategoryRepository) InitTable() error {
	return c.mysqlDB.CreateTable(&model.Category{}).Error
}

func (c *CategoryRepository) FindCategoryByID(cid int64) (category *model.Category, err error) {
	return category, c.mysqlDB.First(&category, cid).Error
}

func (c CategoryRepository) CreateCategory(category *model.Category) (int64, error) {
	return category.ID, c.mysqlDB.Create(category).Error
}

func (c CategoryRepository) DeleteCategoryByID(cid int64) error {
	return c.mysqlDB.Where("id = ?", cid).Delete(&model.Category{}).Error
}

func (c CategoryRepository) UpdateCategory(category *model.Category) error {
	return c.mysqlDB.Model(category).Update(category).Error
}

func (c CategoryRepository) FindAll() (categoryAll []model.Category, err error) {
	return categoryAll, c.mysqlDB.Find(&categoryAll).Error
}

func (c CategoryRepository) FindCategoryByName(name string) (category *model.Category, err error) {
	//category := &model.Category{}
	return category, c.mysqlDB.Where("category_name = ?", name).Find(&category).Error
}

func (c CategoryRepository) FindCategoryByLevel(level uint32) (categorySlice []model.Category, err error) {
	return categorySlice, c.mysqlDB.Where("category_level = ?", level).Find(&categorySlice).Error
}

func (c CategoryRepository) FindCategoryByParent(parentID int64) (categorySlice[]model.Category, err error) {
	return categorySlice, c.mysqlDB.Where("category_parent = ?", parentID).Find(&categorySlice).Error
}
