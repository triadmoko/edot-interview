package services

import (
	"github.com/sirupsen/logrus"
	"github.com/triadmoko/edot-interview/models"
	"github.com/triadmoko/edot-interview/repositorys"
)

type Services interface {
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(productInput models.Product) (*models.Product, error)
	UpdateProduct(productInput models.Product) (*models.Product, error)
	GetProducts() ([]*models.Product, error)
	DeleteProduct(id int) error
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
