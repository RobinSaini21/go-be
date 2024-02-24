package main

import (
	"net/http"

	"begain.com/contollers"
	"begain.com/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}).GET("/check-db", contollers.GetMovies)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
