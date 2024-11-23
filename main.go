package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectDatabase()

	r := gin.Default()

	// Routes
	r.GET("/products", GetProducts)
	r.POST("/products", CreateProduct)
	r.GET("/products/:id", GetProductByID)
	r.PUT("/products/:id", UpdateProduct)
	r.DELETE("/products/:id", DeleteProduct)
	r.POST("/products/:id/image", UploadImage)
	r.GET("/products/:id/image", DownloadImage)

	r.Run()
}
