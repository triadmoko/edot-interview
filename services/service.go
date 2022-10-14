package services

import (
	"github.com/triadmoko/edot-interview/helpers"
	"github.com/triadmoko/edot-interview/repositorys"
)

type Services interface {
}
type services struct {
	repositorys repositorys.Repositorys
	logger      *helpers.StandardLogger
}

func NewService(repositorys repositorys.Repositorys, logger *helpers.StandardLogger) *services {
	return &services{repositorys, logger}
}
