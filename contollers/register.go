package contollers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"begain.com/db"
	"begain.com/jwt"
	"begain.com/structutil"
)

// "john_doe"
type User struct {
	ID                string   `bson:"_id"`
	Username          string   `bson:"username"`
	Email             string   `bson:"email"`
	ReferralBonus     int      `bson:"referralBonus"`
	ReferralCode      string   `bson:"referral_code"`
	ChildrenUsers     []string `bson:"children_user"`
	Password          string   `bson:"password"`
	ReferralRemaining int      `bson:"referral_remaining"`
	IsAdmin           bool     `bson:"is_admin"`
}

func Register(c *gin.Context) {
	var userBody structutil.RequestBody
	if err := c.BindJSON(&userBody); err != nil {
		// If there's an error in binding JSON data, respond with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := db.GetDataBase(c, "referrals")
	collection := database.Collection("users")

	var result User
	filter := bson.D{{"username", userBody.Username}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		fmt.Println("Failed to find document:", err)
		return
	}
	token, err := jwt.GenerateToken(userBody)
	if err != nil {
		println("Not enough")
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "db ok",
		"data":    result,
		"token":   token,
	})
}
