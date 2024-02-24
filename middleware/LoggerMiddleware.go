package middleware

import "github.com/gin-gonic/gin"

func LoggerMiddleware(c *gin.Context) {
	// Before handling the request

	// Call the next handler
	println(c.Request.Method, c.GetHeader("origin"))
	c.Next()

	// After handling the request
}
