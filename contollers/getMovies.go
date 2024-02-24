package contollers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"begain.com/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Document struct {
	ID      string    `json:"_id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	MovieID string    `json:"movie_id"`
	Text    string    `json:"text"`
	Date    time.Time `json:"date"`
}

const MONGODB_ATLAS_URL string = "mongodb+srv://robinsaini2126:MLcB98hSZmTIQTWS@cluster0.x83bzir.mongodb.net/?retryWrites=true&w=majority"

func GetMovies(c *gin.Context) {

	database := db.GetDataBase(c)
	collection := database.Collection("comments")

	queryOptions := options.Find()
	queryOptions.SetLimit(1)

	cursor, err := collection.Find(context.Background(), bson.D{}, queryOptions)
	if err != nil {
		fmt.Println("Failed to find documents:", err)
		return
	}
	defer cursor.Close(context.Background())

	var results []Document
	for cursor.Next(context.Background()) {
		var person Document
		err := cursor.Decode(&person)
		if err != nil {
			fmt.Println("Failed to decode document:", err)
			return
		}
		results = append(results, person)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println("Error during iteration:", err)
		return
	}

	if err != nil {
		fmt.Println("Failed to disconnect from MongoDB:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "db ok",
		"data":    results,
	})
}
