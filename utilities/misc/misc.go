package misc

import (
	"context"

	"genz_server/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckUserExist(db *mongo.Database, user models.Users) (bool, string) {
	var result primitive.M 
	resultErr:= db.Collection("users").FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&result)
	if resultErr != nil {
		return false, ""
	}

	if result["Email"] == "" {
		return false, ""
	}

	return true, "user already exist"
}