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
	r.Static("/static", "./static") // serve static files from the "static" directory

	r.LoadHTMLGlob("views/*.html")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}).GET("/check-db", contollers.GetMovies).POST("/register", contollers.Register).GET("/login-page", func(c *gin.Context) {
		data := gin.H{
			"Name": "Gin User",
		}

		// Render the HTML template with the provided data
		c.HTML(http.StatusOK, "index.html", data)
	}).GET("/api/data", func(c *gin.Context) {
		htmlContent := "<h1>Hello, World!</h1><p>This is HTML content sent in response.</p>"

		// Send HTML content in response
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, htmlContent)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
