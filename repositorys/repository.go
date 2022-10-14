package repositorys

import (
	"database/sql"
	"errors"

	"github.com/triadmoko/edot-interview/models"
	"gorm.io/gorm"
)

type Repositorys interface {
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(product models.Product) (*models.Product, error)
	UpdateProduct(product models.Product) (*models.Product, error)
	DeleteProduct(id string) error
	GetProducts() ([]*models.Product, error)
}
type repositorys struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositorys {
	return &repositorys{db}
}
func (r *repositorys) GetProductByID(id int) (*models.Product, error) {
	var product models.Product

	err := r.db.First(&product, id).Error
	if err == sql.ErrNoRows {
		return nil, errors.New("Product Not Found")
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *repositorys) CreateProduct(product models.Product) (*models.Product, error) {

	err := r.db.Create(&product).Error
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

func (r *repositorys) DeleteProduct(id string) error {
	var product *models.Product
	err := r.db.Where("id = ? ", id).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *repositorys) GetProducts() ([]*models.Product, error) {
	var product []*models.Product
	err := r.db.Find(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
