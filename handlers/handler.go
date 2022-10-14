package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/triadmoko/edot-interview/helpers"
	"github.com/triadmoko/edot-interview/models"
	"github.com/triadmoko/edot-interview/services"
)

type handlers struct {
	services services.Services
}

func NewHandler(services services.Services) *handlers {
	return &handlers{services}
}
func (h *handlers) CreateProduct(c *gin.Context) {
	var req models.Product

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("failed create product", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	product, err := h.services.CreateProduct(req)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		respon := helpers.ResponseApi("failed create product", http.StatusInternalServerError, "Failed", errorMessage)
		c.JSON(http.StatusInternalServerError, respon)
		return
	}

	response := helpers.ResponseApi("success create product", http.StatusCreated, "Success", product)
	c.JSON(http.StatusCreated, response)

}
