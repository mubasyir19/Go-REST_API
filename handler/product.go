package handler

import (
	"net/http"
	"serviceLayer/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	products, err := h.productService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	var productsReponse []product.ProductResponse

	for _, prod := range products {
		productReponse := convertProductResponse(prod)

		productsReponse = append(productsReponse, productReponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productsReponse,
	})
}

func (h *productHandler) GetProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	prod, err := h.productService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	productResponse := convertProductResponse(prod)

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var productRequest product.ProductRequest

	err := c.ShouldBind(&productRequest)
	if err != nil {
		// errorMessages := []string{}
		// for _, e := range err.(validator.ValidationErrors) {
		// 	errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
		// 	errorMessages = append(errorMessages, errorMessage)
		// }

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	product, err := h.productService.Create(productRequest)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    convertProductResponse(product),
		"message": "Successfully added the product",
	})
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var productRequest product.ProductRequest

	err := c.ShouldBind(&productRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	product, err := h.productService.Update(id, productRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    convertProductResponse(product),
		"message": "This book has been updated",
	})

}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	product, err := h.productService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	productResponse := convertProductResponse(product)

	c.JSON(http.StatusOK, gin.H{
		"data":    productResponse,
		"message": "This book has been deleted",
	})

}

func convertProductResponse(prod product.Product) product.ProductResponse {
	return product.ProductResponse{
		ID:          prod.ID,
		Title:       prod.Title,
		Price:       prod.Price,
		Description: prod.Description,
	}
}
