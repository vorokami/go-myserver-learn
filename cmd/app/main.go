package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ping")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run(":8080")
}
