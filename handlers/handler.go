package handlers

import (
	"fmt"
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

// Product
// @Tags Product
// @Summary Create New Product
// @Description Create New Product
// @Accept       json
// @Produce      json
// @Param        product  body      models.Product  true  "Product"
// @Router /product [post]
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

// Product
// @Tags Product
// @Summary Get Product By Product
// @Description Get Product By Product
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Router /product/{id} [get]
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

// Product
// @Tags Product
// @Summary Update Product
// @Description Update Product
// @Accept       json
// @Produce      json
// @Param1        id   path      int  true  "Product ID"
// @Router /product/{id} [put]
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

// Product
// @Tags Product
// @Summary Update Patch Product
// @Description Update Patch Product
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Router /product/{id} [patch]
func (h *handlers) UpdateProductPatch(c *gin.Context) {
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

// Product
// @Tags Product
// @Summary Get All Product
// @Description Get All Product
// @Accept       json
// @Produce      json
// @Router /product [get]
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

// Product
// @Tags Product
// @Summary Delete Product
// @Description Delete Product
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Router /product/{id} [delete]
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

// Category
// @Tags Category
// @Summary Create New Category
// @Description  Create New Category
// @Accept       json
// @Produce      json
// @Param        Category  body      models.Category  true  "Category"
// @Router /category [post]
func (h *handlers) CreateCategory(c *gin.Context) {
	var req models.Category

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("failed create category", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	category, err := h.services.CreateCategory(req)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		respon := helpers.ResponseApi("failed create category", http.StatusInternalServerError, "Failed", errorMessage)
		c.JSON(http.StatusInternalServerError, respon)
		return
	}

	response := helpers.ResponseApi("success create category", http.StatusCreated, "Success", category)
	c.JSON(http.StatusCreated, response)
}

// Category
// @Tags Category
// @Summary Get Category By Category
// @Description Get Category By Category
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Router /category/{id} [get]
func (h *handlers) GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("id not int", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	Category, err := h.services.GetCategoryByID(id)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		respon := helpers.ResponseApi("failed get Category", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	response := helpers.ResponseApi("success get Category", http.StatusOK, "Success", Category)
	c.JSON(http.StatusOK, response)
}

// Category
// @Tags Category
// @Summary Update Category
// @Description Update Category
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Router /category/{id} [put]
func (h *handlers) UpdateCategory(c *gin.Context) {
	var req models.Category
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
		respon := helpers.ResponseApi("failed update Category", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	category, err := h.services.UpdateCategory(req)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		respon := helpers.ResponseApi("failed update Category", http.StatusInternalServerError, "Failed", errorMessage)
		c.JSON(http.StatusInternalServerError, respon)
		return
	}
	fmt.Println(category)
	response := helpers.ResponseApi("success update Category", http.StatusOK, "Success", category)
	c.JSON(http.StatusOK, response)

}

// Category
// @Tags Category
// @Summary Delete Category
// @Description Delete Category
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Router /category/{id} [patch]
func (h *handlers) UpdateCategoryPatch(c *gin.Context) {
	var req models.Category
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
		respon := helpers.ResponseApi("failed update Category", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	category, err := h.services.UpdateCategory(req)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		respon := helpers.ResponseApi("failed update Category", http.StatusInternalServerError, "Failed", errorMessage)
		c.JSON(http.StatusInternalServerError, respon)
		return
	}
	fmt.Println(category)
	response := helpers.ResponseApi("success update Category", http.StatusOK, "Success", category)
	c.JSON(http.StatusOK, response)
}

// Category
// @Tags Category
// @Summary Get All Category
// @Description Get All Category
// @Accept       json
// @Produce      json
// @Router /category [get]
func (h *handlers) GetCategorys(c *gin.Context) {

	Categorys, err := h.services.GetCategorys()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		respon := helpers.ResponseApi("failed get all Category", http.StatusInternalServerError, "Failed", errorMessage)
		c.JSON(http.StatusInternalServerError, respon)
		return
	}

	response := helpers.ResponseApi("success get all Category", http.StatusOK, "Success", Categorys)
	c.JSON(http.StatusOK, response)

}

// Category
// @Tags Category
// @Summary Delete Category
// @Description Delete Category
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Router /category/{id} [delete]
func (h *handlers) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("id not int", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	err = h.services.DeleteCategory(id)
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("failed delete Category", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	response := helpers.ResponseApi("success delete Category", http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)
}
