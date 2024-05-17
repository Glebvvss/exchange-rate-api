package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Glebvvss/exchange-rate-api/internal"
	"github.com/Glebvvss/exchange-rate-api/internal/log"
)

func GetRate(c *gin.Context) {
	rate, err := internal.GetUsdRate()

	if err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"error": "Wrong email format"})
		return
	}

	c.JSON(http.StatusOK, rate)
}
