package controllers

import (

	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mahathirrizky/go-vue/database"
	"github.com/mahathirrizky/go-vue/models"
)

// CreateContent creates a new content entry
func CreateContent(c *gin.Context) {
	var input models.Content
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Content created successfully", "content": input})
}

// GetContent retrieves content by page name, section, and key
func GetContent(c *gin.Context) {
	pageName := c.Query("page_name")
	section := c.Query("section")
	key := c.Query("key")

	var content models.Content
	query := database.DB.Where("page_name = ?", pageName)

	if section != "" {
		query = query.Where("section = ?", section)
	}
	if key != "" {
		query = query.Where("key = ?", key)
	}

	if err := query.First(&content).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Content not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"content": content})
}

// UpdateContent updates an existing content entry
func UpdateContent(c *gin.Context) {
	id := c.Param("id")

	var content models.Content
	if err := database.DB.First(&content, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Content not found"})
		return
	}

	var input models.Content
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the updated key is a file URL (photo or CV) and if the value has changed
	if (input.Key == "photo_url" || input.Key == "cv_download_url") && input.Value != content.Value {
		// Extract the filename from the old URL
		oldFilename := filepath.Base(content.Value)
		oldFilePath := filepath.Join("uploads", oldFilename)

		// Check if the old file exists and is not the default pp.png
		// We don't want to delete the default image
		if oldFilename != "pp.png" {
			if _, err := os.Stat(oldFilePath); err == nil {
				// Delete the old file
				if err := os.Remove(oldFilePath); err != nil {
					// Log the error but don't stop the update process
					// as deleting the old file is not critical for the update itself
					c.Error(err).SetType(gin.ErrorTypePrivate)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old photo"})
					return
				}
			}
		}
	}

	database.DB.Model(&content).Updates(input)
	c.JSON(http.StatusOK, gin.H{"message": "Content updated successfully", "content": content})
}

// DeleteContent deletes a content entry
func DeleteContent(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Content{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Content deleted successfully"})
}

// GetAllContent retrieves all content entries
func GetAllContent(c *gin.Context) {
	var contents []models.Content
	if err := database.DB.Find(&contents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve content"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"contents": contents})
}

// UploadFile handles file uploads
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	// Generate a new UUID for the filename
	newUUID, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
		return
	}

	// Create a new filename with the original extension
	ext := filepath.Ext(file.Filename)
	newFilename := newUUID.String() + ext

	// Define the path to save the file
	// The path is relative to the project root: ui/public/uploads
	uploadPath := "ui/public/uploads"

	// Create the uploads directory if it doesn't exist
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// Save the file
	dst := filepath.Join(uploadPath, newFilename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
		return
	}

	// Return the URL of the uploaded file
	fileURL := "/uploads/" + newFilename
	c.JSON(http.StatusOK, gin.H{"file_path": fileURL, "message": "File uploaded successfully"})
}
