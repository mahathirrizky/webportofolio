package main

import (
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/mahathirrizky/go-vue/database"
    "github.com/mahathirrizky/go-vue/routes"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Connect to database
    database.ConnectDatabase()
    database.Seed()

    // Set Gin to production mode
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()

    // CORS configuration
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusOK)
            return
        }
        c.Next()
    })

    // API routes (register these first)
    routes.SetupRoutes(r)

    // Serve static files from ui/dist
    r.Static("/assets", filepath.Join(os.Getenv("PWD"), "ui", "dist", "assets"))
    r.Static("/uploads", "./uploads")
    // Serve specific files in ui/dist (avoid catch-all)
    r.StaticFS("/static", http.Dir(filepath.Join(os.Getenv("PWD"), "ui", "dist")))

    // Serve Vue app for any unmatched routes
    r.NoRoute(func(c *gin.Context) {
        c.File(filepath.Join(os.Getenv("PWD"), "ui", "dist", "index.html"))
    })

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Server listening on :%s\n", port)
    log.Fatal(r.Run(":" + port))
}