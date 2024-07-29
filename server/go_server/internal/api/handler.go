package api

import (
	"fmt"
	"go_server/internal/model"
	"go_server/internal/service"
	"go_server/pkg/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(c *gin.Context) {
	if err := service.PingDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", "message": "Database connection failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok", "message": "Server is running and database is connected",
	})
}

func HandleClipboardUpload(c *gin.Context) {
	var clip model.Clip
	if err := c.ShouldBindJSON(&clip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.SaveClipboard(clip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save clipboard content",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": clip.CreatedAt,
	})
}

func HandleUserSignUp(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.ErrorLogger.Printf("Invalid user data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.SaveUser(c, user)

	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			logger.InfoLogger.Printf("Attempt to register existing user: %s", user.Name)
			c.JSON(http.StatusConflict, gin.H{
				"error": fmt.Sprintf("User %s already exists", user.Name),
			})
		} else {
			logger.ErrorLogger.Printf("Failed to save user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to process user registration",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("New user created: %s", user.Name),
	})
}

func HandleDeviceSignUp(c *gin.Context) {
	var device model.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		logger.ErrorLogger.Printf("Invalid device data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.SaveDevice(c, device)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			logger.InfoLogger.Printf("Attempt to register existing device: %s", device.Name)
			c.JSON(http.StatusConflict, gin.H{
				"error": fmt.Sprintf("Device %s already exists", device.Name),
			})
		} else {
			logger.ErrorLogger.Printf("Failed to save device: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to process device registration",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("New device created: %s", device.Name),
	})
}
