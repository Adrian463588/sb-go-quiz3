package controllers

import (
	"errors"
	"net/http"

	"go-bookstore-api/config"
	"go-bookstore-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func GetCategories(c *gin.Context) {
	var categories []models.Category
	config.DB.Find(&categories)
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func GetCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := config.DB.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func CreateCategory(c *gin.Context) {
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{Name: input.Name}
	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan kategori"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": category})
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := config.DB.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori yang akan dihapus tidak ditemukan"}) // Validasi
		return
	}

	config.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil dihapus"})
}

func GetBooksByCategory(c *gin.Context) {
	var books []models.Book
	id := c.Param("id")

	// Preload("Category") untuk menyertakan data kategori di setiap buku
	if err := config.DB.Preload("Category").Where("category_id = ?", id).Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}