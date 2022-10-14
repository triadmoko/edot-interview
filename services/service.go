package services

import (
	"github.com/sirupsen/logrus"
	"github.com/triadmoko/edot-interview/models"
	"github.com/triadmoko/edot-interview/repositorys"
)

type Services interface {
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(productInput models.Product) (*models.Product, error)
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
