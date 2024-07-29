package main

import (
	"go_server/internal/api"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 设置路由
	r.POST("/clip", api.HandleClipboardUpload)
	r.GET("/health", api.HandleHealthCheck)
	r.POST("/user")

	// 启动服务器
	log.Println("Server is running on :9527")
	if err := r.Run(":9527"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
