package repositorys

import (
	"errors"

	"github.com/triadmoko/edot-interview/models"
	"gorm.io/gorm"
)

type Repositorys interface {
	// product
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(product models.Product) (*models.Product, error)
	UpdateProduct(product models.Product) (*models.Product, error)
	DeleteProduct(id int) error
	GetProducts() ([]*models.Product, error)

	// category
	GetCategoryByID(id int) (*models.Category, error)
	CreateCategory(category models.Category) (*models.Category, error)
	UpdateCategory(category models.Category) (*models.Category, error)
	DeleteCategory(id int) error
	GetCategorys() ([]*models.Category, error)
}
type repositorys struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositorys {
	return &repositorys{db}
}
func (r *repositorys) GetProductByID(id int) (*models.Product, error) {
	var product models.Product

	err := r.db.Preload("Category").First(&product, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Product Not Found")
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *repositorys) CreateProduct(product models.Product) (*models.Product, error) {

	err := r.db.Preload("Category").Create(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *repositorys) UpdateProduct(product models.Product) (*models.Product, error) {
	err := r.db.Where("id = ? ", product.ID).Save(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repositorys) DeleteProduct(id int) error {
	var product *models.Product
	err := r.db.Where("id = ? ", id).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *repositorys) GetProducts() ([]*models.Product, error) {
	var product []*models.Product
	err := r.db.Preload("Category").Find(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repositorys) GetCategoryByID(id int) (*models.Category, error) {
	var category models.Category

	err := r.db.First(&category, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Category Not Found")
	}
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *repositorys) CreateCategory(category models.Category) (*models.Category, error) {

	err := r.db.Create(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *repositorys) UpdateCategory(category models.Category) (*models.Category, error) {
	err := r.db.Where("id = ? ", category.ID).Save(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *repositorys) DeleteCategory(id int) error {
	var category *models.Category
	err := r.db.Where("id = ? ", id).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repositorys) GetCategorys() ([]*models.Category, error) {
	var category []*models.Category
	err := r.db.Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
