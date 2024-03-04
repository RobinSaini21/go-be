package contollers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"begain.com/structutil"
	"github.com/gin-gonic/gin"
)

// "john_doe"

func Product(c *gin.Context) {
	id := c.Param("id")
	requestUrl := fmt.Sprintf("https://fakestoreapi.com/products/%s", id)
	resp, err := http.Get(requestUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into a slice of Post structs
	var Product structutil.Product
	if err := json.NewDecoder(resp.Body).Decode(&Product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "product.html", gin.H{
		"Product": Product,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "db ok",
	// 	"data":    Products,
	// })
}
