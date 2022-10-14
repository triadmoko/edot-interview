package handlers

import (
	"net/http"
	"strconv"

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

func (h *handlers) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("id not int", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	product, err := h.services.GetProductByID(id)
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("failed get product", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	response := helpers.ResponseApi("success get product", http.StatusOK, "Success", product)
	c.JSON(http.StatusOK, response)
}

func (h *handlers) UpdateProduct(c *gin.Context) {
	var req models.Product
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("id not int", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	req.ID = uint(id)
	err = c.ShouldBindJSON(&req)
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("failed update product", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	product, err := h.services.UpdateProduct(req)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		respon := helpers.ResponseApi("failed update product", http.StatusInternalServerError, "Failed", errorMessage)
		c.JSON(http.StatusInternalServerError, respon)
		return
	}

	response := helpers.ResponseApi("success update product", http.StatusOK, "Success", product)
	c.JSON(http.StatusOK, response)

}

func (h *handlers) GetProducts(c *gin.Context) {

	products, err := h.services.GetProducts()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		respon := helpers.ResponseApi("failed get all product", http.StatusInternalServerError, "Failed", errorMessage)
		c.JSON(http.StatusInternalServerError, respon)
		return
	}

	response := helpers.ResponseApi("success get all product", http.StatusOK, "Success", products)
	c.JSON(http.StatusOK, response)

}

func (h *handlers) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("id not int", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	err = h.services.DeleteProduct(id)
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("failed delete product", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	response := helpers.ResponseApi("success delete product", http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)
}
