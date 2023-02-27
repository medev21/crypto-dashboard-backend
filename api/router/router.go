package router

import (
	"medev21/crypto-dashboard-backend/api/handlers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	eng := gin.Default()

	eng.GET("/healthcheck", handlers.GetHealthCheck)
	eng.GET("/users", handlers.GetUsers)

	return eng
}
