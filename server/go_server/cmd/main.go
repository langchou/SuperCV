package main

import (
	"go_server/internal/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	logDir := "./logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	r := gin.Default()

	// 设置路由
	r.POST("/clip", api.HandleClipboardUpload)
	r.GET("/health", api.HandleHealthCheck)

	r.POST("/user/signup", api.HandleUserSignUp)
	r.POST("/device/signup", api.HandleDeviceSignUp)

	// 启动服务器
	log.Println("Server is running on :9527")
	if err := r.Run(":9527"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
