package api

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"

	"github.com/Glebvvss/exchange-rate-api/internal"
	"github.com/Glebvvss/exchange-rate-api/internal/log"
)

func PostSubscribe(c *gin.Context) {
	type FormParams struct {
		Email string `form:"email" binding:"required"`
	}

	type JsonParams struct {
		Email string `json:"email" binding:"required"`
	}

	var email string
	var jsonParams JsonParams
	var formParams FormParams
	if err := c.BindJSON(&jsonParams); err == nil {
		email = jsonParams.Email
	} else if err := c.Bind(&formParams); err == nil {
		email = formParams.Email
	} else {
		c.JSON(400, gin.H{"error": "Invalid data provided"})
		return
	}

	if err := checkmail.ValidateFormat(email); err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"error": "Wrong email format"})
		return
	}

	if err := internal.SubscribeEmail(email); err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"error": "Somethind went wrong"})
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
