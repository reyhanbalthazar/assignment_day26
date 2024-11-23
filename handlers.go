package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []Product
	DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var input Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&product).Updates(input)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func UploadImage(c *gin.Context) {
	id := c.Param("id")
	file, _ := c.FormFile("file")
	path := "./images/" + id + ".jpg"
	c.SaveUploadedFile(file, path)
	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded"})
}

func DownloadImage(c *gin.Context) {
	id := c.Param("id")
	path := "./images/" + id + ".jpg"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.File(path)
}
