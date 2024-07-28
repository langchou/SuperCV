package main

import (
	"go_server/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api.SetupRoutes(router)

	router.Run(":9527")
}
