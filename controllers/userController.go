package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shrisuradkar/RestAPI-MonogoDB-Golang/database"
	"github.com/shrisuradkar/RestAPI-MonogoDB-Golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("List all Users")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var users []primitive.M
		cursor, err := collection.Find(ctx, bson.D{{}})
		defer cancel()
		if err != nil {
			log.Fatal(err)
		}
		for cursor.Next(ctx) {
			var user bson.M
			err := cursor.Decode(&user)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
		defer cursor.Close(ctx)
		c.JSON(http.StatusOK, users)
	}
}

func GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Get user by id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		userId := c.Param("id")

		objId, _ := primitive.ObjectIDFromHex(userId)

		var user models.User
		filter := bson.M{"_id": objId}
		err := collection.FindOne(ctx, filter).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Creating User")

		var user models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		//check is user already present or not cheking with email id
		emailFilter := bson.M{"email": user.Email}
		countE, err := collection.CountDocuments(ctx, emailFilter)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking email "})
			return
		}
		//check is user already present or not cheking with phone number
		phoneFilter := bson.M{"phone": user.Phone}
		countP, err := collection.CountDocuments(ctx, phoneFilter)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while checking phone "})
			return
		}

		if countE > 0 && countP > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "user with same emailID or phone number already present"})
			return
		}
		newUser := models.User{
			ID:        primitive.NewObjectID(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Age:       user.Age,
			Email:     user.Email,
			Gender:    user.Gender,
			Phone:     user.Phone,
		}
		result, insertErr := collection.InsertOne(ctx, newUser)
		defer cancel()
		if insertErr != nil {
			log.Fatal(insertErr.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not created"})
			return
		}
		fmt.Println("result", result.InsertedID)
		c.JSON(http.StatusOK, gin.H{"success": "user created successfully"})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Updating existing user")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		userId := c.Param("id")
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		objId, _ := primitive.ObjectIDFromHex(userId)

		//check user already present or not
		checkFilter := bson.M{"_id": objId}
		count, err := collection.CountDocuments(ctx, checkFilter)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not exist"})
			return
		}
		filter := bson.M{"_id": objId}
		update := bson.M{"$set": bson.M{"watched": true}}
		result, err := collection.UpdateOne(ctx, filter, update)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if result.ModifiedCount > 0 {
			c.JSON(http.StatusOK, gin.H{"sucess": "User Updated Successfully"})
		}
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		userID := c.Param("id")
		objId, _ := primitive.ObjectIDFromHex(userID)

		//check is user present of not
		checkFilter := bson.M{"_id": objId}
		count, err := collection.CountDocuments(ctx, checkFilter)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
			return
		}
		filter := bson.M{"_id": objId}
		result, deleteErr := collection.DeleteOne(ctx, filter)
		defer cancel()
		if deleteErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if result.DeletedCount > 0 {
			c.JSON(http.StatusOK, gin.H{"success": "user deleted successfully"})
		}

	}
}
