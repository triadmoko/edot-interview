package handlers

import (
	"github.com/triadmoko/edot-interview/services"
)

type handlers struct {
	services services.Services
}

func NewHandler(services services.Services) *handlers {
	return &handlers{services}
}
