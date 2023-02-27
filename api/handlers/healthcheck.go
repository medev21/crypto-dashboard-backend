package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "everything is good homie!",
	})
}
