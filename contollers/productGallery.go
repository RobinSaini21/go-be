package contollers

import (
	"encoding/json"
	"net/http"

	"begain.com/structutil"
	"github.com/gin-gonic/gin"
)

// "john_doe"

func ProductGallery(c *gin.Context) {

	resp, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into a slice of Post structs
	var Products []structutil.Product
	if err := json.NewDecoder(resp.Body).Decode(&Products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "products.html", gin.H{
		"Products": Products,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "db ok",
	// 	"data":    Products,
	// })
}
