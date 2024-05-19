package main

import "github.com/gin-gonic/gin"

func main() {
	gin.Default().GET("/hello", func(c *gin.Context) { c.String(200, "hello, this is vinay surya") }).Run(":8080")
}
