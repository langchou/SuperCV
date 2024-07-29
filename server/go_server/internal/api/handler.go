package api

import (
	"go_server/internal/model"
	"go_server/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		"message": "Clipboard content saved successfully",
	})
}

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

func HandleUserSignIn(c *gin.Context) {

}
