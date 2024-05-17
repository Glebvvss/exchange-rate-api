package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"

	"github.com/Glebvvss/exchange-rate-api/api"
	"github.com/Glebvvss/exchange-rate-api/crontask"
	"github.com/Glebvvss/exchange-rate-api/internal"
)

func init() {
	time.Sleep(10 * time.Second)
	internal.Migrate()
}

func main() {
	c := cron.New()
	c.AddFunc("0 5 * * *", crontask.SendEmails)
	c.Start()

	router := gin.Default()
	apiGroup := router.Group("/api")
	apiGroup.GET("/rate", api.GetRate)
	apiGroup.POST("/subscribe", api.PostSubscribe)
	if err := router.Run(); err != nil {
		panic(err)
	}
}
