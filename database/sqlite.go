package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/mahathirrizky/go-vue/models"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("people.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate models here
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Content{})
	DB.AutoMigrate(&models.Message{})

	SeedAdminUser() // Still seed admin user
}

func SeedAdminUser() {
	var user models.User
	// Check if an admin user already exists
	result := DB.Where("is_admin = ?", true).First(&user)
	if result.Error == nil {
		log.Println("Admin user already exists, skipping seeding.")
		return // Admin user already exists
	}

	// Create a default admin user if none exists
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("adminpassword"), bcrypt.DefaultCost) // Use a strong default password
	if err != nil {
		log.Fatal("Failed to hash admin password:", err)
	}

	adminUser := models.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: string(hashedPassword),
		IsAdmin:  true,
	}

	if err := DB.Create(&adminUser).Error; err != nil {
		log.Fatal("Failed to create admin user:", err)
	}

	log.Println("Default admin user created: admin@example.com with password 'adminpassword'")
}