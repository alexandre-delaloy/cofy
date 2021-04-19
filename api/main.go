package main

import (
	"net/http"

	"github.com/blyndusk/cofy/api/database"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	setupServer()
}

func setupServer() *gin.Engine {

	cfg := database.Config{}
	log.Info(cfg)
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	err := database.Init(cfg)

	if err != nil {
		log.Fatal("fail to connect db", err)
	}
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3003") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return r
}
