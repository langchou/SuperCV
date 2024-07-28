/*
处理所有的HTTP请求，并调用相应的服务逻辑
*/

package api

import (
	"go_server/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/clip", getClip)
}

func getClip(c *gin.Context) {
	svc := service.NewClipboardService()
	clip, err := svc.GetLatestClip()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clip)
}
