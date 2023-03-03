package handlers

import (
	geminiapi "medev21/gogemini-api"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGeminiTicker(c *gin.Context) {

	gem := geminiapi.New(false, "api key", "api secret")

	res, err := gem.GetCoinTicker("btcusd")

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetCryptoSymbols(c *gin.Context) {
	gem := geminiapi.New(true, "api key", "api secret")

	res, err := gem.GetSymbols()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func CheckGeminiEnv(c *gin.Context) {
	gem := geminiapi.New(false, "api key", "api secret")

	res := gem.GetGeminiEnv()

	c.JSON(http.StatusOK, res)
}
