package main

import (
	"fmt"
	"go-bookstore-api/config"
	"go-bookstore-api/models"
	"go-bookstore-api/routes"
	"os" // <-- Impor package 'os'

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Coba muat file .env, tapi jangan hentikan program jika tidak ada.
	godotenv.Load()

	// Koneksi ke database
	config.ConnectDatabase()

	// AutoMigrate untuk membuat tabel sesuai model
	config.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Book{})

	// Inisialisasi Gin router
	r := gin.Default()

	// Setup semua rute
	routes.SetupRouter(r)

	// Menggunakan port dinamis yang diberikan oleh Railway, atau 8080 untuk lokal
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port default untuk pengembangan lokal
	}

	fmt.Println("Starting server on port " + port)
	// Jalankan server di port yang sudah ditentukan
	r.Run(":" + port)
}