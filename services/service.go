package services

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/edot-interview/models"
	"github.com/triadmoko/edot-interview/repositorys"
	"gorm.io/gorm"
)

type Services interface {
	// product
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(productInput models.Product) (*models.Product, error)
	UpdateProduct(productInput models.Product) (*models.Product, error)
	GetProducts() ([]*models.Product, error)
	DeleteProduct(id int) error

	// category
	CreateCategory(inputCategory models.Category) (*models.Category, error)
	GetCategoryByID(id int) (*models.Category, error)
	UpdateCategory(categoryInput models.Category) (*models.Category, error)
	GetCategorys() ([]*models.Category, error)
	DeleteCategory(id int) error
}
type services struct {
	repositorys repositorys.Repositorys
	logger      *logrus.Logger
}

func NewService(repositorys repositorys.Repositorys, logger *logrus.Logger) *services {
	return &services{repositorys, logger}
}

func (s *services) GetProductByID(id int) (*models.Product, error) {
	product, err := s.repositorys.GetProductByID(id)
	if err != nil {
		s.logger.Error("failed get product by id ", err.Error())
		return nil, err
	}
	return product, nil
}

func (s *services) CreateProduct(productInput models.Product) (*models.Product, error) {
	category, err := s.repositorys.GetCategoryByID(productInput.CategoryID)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error("failed get category ", err.Error())
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}
	product, err := s.repositorys.CreateProduct(productInput)
	if err != nil {
		s.logger.Error("failed create product ", err.Error())
		return nil, err
	}
	return product, nil
}
func (s *services) UpdateProduct(productInput models.Product) (*models.Product, error) {
	product, err := s.repositorys.UpdateProduct(productInput)
	if err != nil {
		s.logger.Error("failed update product ", err.Error())
		return nil, err
	}
	return product, nil
}
func (s *services) GetProducts() ([]*models.Product, error) {
	products, err := s.repositorys.GetProducts()
	if err != nil {
		s.logger.Error("failed get all product ", err.Error())
		return nil, err
	}
	return products, nil
}
func (s *services) DeleteProduct(id int) error {
	err := s.repositorys.DeleteProduct(id)
	if err != nil {
		s.logger.Error("failed delete product ", err.Error())
		return err
	}
	return nil
}

func (s *services) CreateCategory(inputCategory models.Category) (*models.Category, error) {
	category, err := s.repositorys.CreateCategory(inputCategory)
	if err != nil {
		s.logger.Error("failed create category ", err.Error())
		return nil, err
	}
	return category, nil
}

func (s *services) GetCategoryByID(id int) (*models.Category, error) {
	category, err := s.repositorys.GetCategoryByID(id)
	if err != nil {
		s.logger.Error("failed get Category by id ", err.Error())
		return nil, err
	}
	return category, nil
}

func (s *services) UpdateCategory(CategoryInput models.Category) (*models.Category, error) {
	category, err := s.repositorys.UpdateCategory(CategoryInput)
	if err != nil {
		s.logger.Error("failed update category ", err.Error())
		return nil, err
	}
	return category, nil
}
func (s *services) GetCategorys() ([]*models.Category, error) {
	categorys, err := s.repositorys.GetCategorys()
	if err != nil {
		s.logger.Error("failed get all Category ", err.Error())
		return nil, err
	}
	return categorys, nil
}
func (s *services) DeleteCategory(id int) error {
	err := s.repositorys.DeleteCategory(id)
	if err != nil {
		s.logger.Error("failed delete Category ", err.Error())
		return err
	}
	return nil
}
