package main

import (
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"/ping":                  "endpoint that response with a json (pong).",
			"/pingTime":              "endpoint that response with the server time.",
			"/getServerHardwardData": "Get number of cores, golang version, etc.",
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

	r.GET("/getServerHardwardData", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"You are using ":        runtime.Compiler,
			"on a":                  runtime.GOARCH,
			"Using Go version":      runtime.Version(),
			"Number of CPUs:":       runtime.NumCPU(),
			"Number of Goroutines:": runtime.NumGoroutine(),
		})
	})

	if os.Getenv("PORT") != "" {
		// Heroku add a env variable called PORT, if exist we will use it
		r.Run("0.0.0.0:" + os.Getenv("PORT"))
	} else {
		// If is running on localhost (our computer), no PORT env variable
		r.Run("0.0.0.0:8080")
	}
}
