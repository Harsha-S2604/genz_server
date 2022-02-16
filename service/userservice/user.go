package userservice

import (
	"net/http"
	"log"
	"context"

	"genz_server/models"
	"genz_server/utilities/validations"
	"genz_server/utilities/hashing"
	"genz_server/utilities/misc"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserHandler(db *mongo.Database) gin.HandlerFunc {

	addUser := func(ctx *gin.Context) {
		var user models.Users
		/*
			1. bind request body with user object declared above
			2. validate user email by calling ValidateUserEmail from package validations
			3. check if user exists
				if exist 
					don't add user to db
			4. hash user password
			5. add user to database
		*/

		ctx.ShouldBindJSON(&user)
		isValidUserData, msg := validations.ValidateUserData(user.Email, user.Password)
		if !isValidUserData {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": msg,
			})
			return
		}
		isUserExist, msg := misc.CheckUserExist(db, user)
		if isUserExist {
			ctx.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": msg,
			})
			return
		}
		hashedPassword := hashing.HashUserPassword(user.Password)
		user.Password = hashedPassword
		insertResult, insertResultErr := db.Collection("users").InsertOne(ctx, user)
		if insertResultErr != nil {
			log.Println("add user error:", insertResultErr.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Sorry, something went wrong. our team is working on it. Please try again later",
			})
			return
		}
		registeredId := insertResult.InsertedID.(primitive.ObjectID)
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"id": registeredId,
			"message": "User registration successful. Please sign-in to continue.",
		})
	}

	return gin.HandlerFunc(addUser)

}

func UserLoginHandler(db *mongo.Database) gin.HandlerFunc {

	userLogin := func(ctx *gin.Context) {
		var userLoginObj models.UserLogin
		ctx.ShouldBindJSON(&userLoginObj)
		isValidUserData, msg := validations.ValidateUserData(userLoginObj.Email, userLoginObj.Password)
		if !isValidUserData {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": msg,
			})
			return
		}
		hashedPassword := hashing.HashUserPassword(userLoginObj.Password)
		
		var result primitive.M
		resultErr := db.Collection("users").FindOne(context.TODO(), bson.M{"email": userLoginObj.Email}).Decode(&result)
		if resultErr != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Looks like you have not registered. Please register to continue.",
			})
			return
		}
		email, ok := result["email"]
		if !ok || email == "" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Looks like you have not registered. Please register to continue.",
			})
			return
		}
		
		password := result["password"]
		if email != userLoginObj.Email || password != hashedPassword {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "invalid email or password",
			})
			return
		}

		delete(result, "password")
		delete(result, "social_profile")
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "login successful",
			"userData": result,
		})
	}

	return gin.HandlerFunc(userLogin)
}

func GetUserByIdHandler(db *mongo.Database) gin.HandlerFunc {

	getUserById := func (ctx *gin.Context) {
		userIdStr := ctx.Params.ByName("id")
		objectId, objectIdErr := primitive.ObjectIDFromHex(userIdStr)
		if objectIdErr != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "invaid id",
			})
			return
		}
		var user primitive.M
		_ = db.Collection("users").FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user)
		if email, ok := user["email"]; !ok || email == "" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "user not found",
			})
			return
		}

		delete(user, "password")
		ctx.JSON(http.StatusFound, gin.H{
			"success": true,
			"message": "user found",
			"userData": user,
		})
	}
	return gin.HandlerFunc(getUserById)
}