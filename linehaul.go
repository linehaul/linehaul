package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Hello World\n")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}
