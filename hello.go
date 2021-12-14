package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":3000")
}
