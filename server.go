package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"/ping": "endpoint that response with a json (pong).",
			"/pingTime": "endpoint that response with the server time.",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/pingTime", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
