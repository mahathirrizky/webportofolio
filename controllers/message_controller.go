package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahathirrizky/go-vue/database"
	"github.com/mahathirrizky/go-vue/models"
)

// CreateMessage creates a new message
func CreateMessage(c *gin.Context) {
	// Honeypot check
	if c.PostForm("website") != "" {
		c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
		return
	}

	var input models.Message
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

// GetAllMessages retrieves all messages
func GetAllMessages(c *gin.Context) {
	var messages []models.Message
	if err := database.DB.Order("created_at desc").Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

// DeleteMessage deletes a message
func DeleteMessage(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Message{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}
