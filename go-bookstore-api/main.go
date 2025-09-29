package main

import (
	"go-bookstore-api/config"
	"go-bookstore-api/models"
	"go-bookstore-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables dari .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Koneksi ke database
	config.ConnectDatabase()

	// AutoMigrate untuk membuat tabel sesuai model
	config.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Book{})

	// Inisialisasi Gin router
	r := gin.Default()

	// Setup semua rute
	routes.SetupRouter(r)

	// Jalankan server
	r.Run()
}