package main

import (
	"weather-api/internal/config"
	"weather-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	r := gin.Default()
	routes.InitRoutes(r)
	r.Run(":" + config.Port)
}
