package controllers

import (
	"errors"
	"net/http"

	"go-bookstore-api/config"
	"go-bookstore-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int    `json:"release_year" binding:"required,min=1980,max=2024"` // Validasi
	Price       int    `json:"price" binding:"required"`
	TotalPage   int    `json:"total_page" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Preload("Category").Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func GetBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := config.DB.Preload("Category").First(&book, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context) {
	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Logika konversi untuk 'thickness'
	var thickness string
	if input.TotalPage > 100 {
		thickness = "tebal"
	} else {
		thickness = "tipis"
	}

	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		Thickness:   thickness, // Hasil konversi
		CategoryID:  input.CategoryID,
	}

	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan buku"})
		return
	}

	// Load kategori untuk ditampilkan di response
	config.DB.Preload("Category").First(&book, book.ID)
	c.JSON(http.StatusCreated, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := config.DB.First(&book, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku yang akan dihapus tidak ditemukan"}) // Validasi
		return
	}

	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}