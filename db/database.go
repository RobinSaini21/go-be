package db

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_ATLAS_URL string = "mongodb+srv://robinsaini2126:MLcB98hSZmTIQTWS@cluster0.x83bzir.mongodb.net/?retryWrites=true&w=majority"

func GetDataBase(c *gin.Context, databaseName string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(MONGODB_ATLAS_URL)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return nil
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
		return nil
	}

	fmt.Println("Connected to MongoDB!")

	// No need to disconnect here as the connection will be closed when the client goes out of scope.

	fmt.Println("Disconnected from MongoDB!")

	database := client.Database(databaseName)

	return database
}
