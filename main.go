package main

import (
	"fmt"
	"log"
	"serviceLayer/handler"
	"serviceLayer/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Database connection
	dsn := "root:@tcp(127.0.0.1:3306)/service_layer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection Failed")
	}
	fmt.Println("DB connection Success")

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	// Auto Migration Database from Entity
	db.AutoMigrate(product.Product{})

	// Gin Declaration
	router := gin.Default()

	// versioning API
	// v1 := router.Group("/v1")

	// Router Declaration
	router.GET("/products", productHandler.GetProducts)
	router.GET("/products/:id", productHandler.GetProduct)
	router.POST("/products", productHandler.CreateProduct)
	router.PUT("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	// For running Server
	router.Run()
}
