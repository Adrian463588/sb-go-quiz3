package routes

import (
	"go-bookstore-api/controllers"
	"go-bookstore-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Rute untuk user (tidak perlu autentikasi untuk register/login)
		users := api.Group("/users")
		{
			users.POST("/register", controllers.Register)
			users.POST("/login", controllers.Login) // Sesuai dokumen 
		}

		// Rute untuk kategori (dilindungi middleware) [cite: 11]
		categories := api.Group("/categories", middlewares.AuthMiddleware())
		{
			categories.GET("/", controllers.GetCategories)
			categories.POST("/", controllers.CreateCategory)
			categories.GET("/:id", controllers.GetCategory)
			categories.DELETE("/:id", controllers.DeleteCategory)
			categories.GET("/:id/books", controllers.GetBooksByCategory)
		}

		// Rute untuk buku (dilindungi middleware) 
		books := api.Group("/books", middlewares.AuthMiddleware())
		{
			books.GET("/", controllers.GetBooks)
			books.POST("/", controllers.CreateBook)
			books.GET("/:id", controllers.GetBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
}