package main

import (
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

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
