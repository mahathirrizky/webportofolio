package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mahathirrizky/go-vue/controllers"
	"github.com/mahathirrizky/go-vue/middleware"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.LoginUser)
		}

		// Public routes
		api.GET("/content", controllers.GetAllContent) // Public endpoint to get all content
		api.POST("/messages", controllers.CreateMessage)   // Public endpoint to create a message

		// Admin-only routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminAuthMiddleware())
		{
			admin.GET("/content", controllers.GetAllContent)
			admin.POST("/content", controllers.CreateContent)
			admin.PUT("/content/:id", controllers.UpdateContent)
			admin.DELETE("/content/:id", controllers.DeleteContent)
			admin.POST("/upload", controllers.UploadFile) // Route for file uploads

			// Message routes
			admin.GET("/messages", controllers.GetAllMessages)
			admin.DELETE("/messages/:id", controllers.DeleteMessage)
		}
	}
}