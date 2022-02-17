package db

import (
	"os"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func ConnectDB() *mongo.Database {
	username, password := os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
	if username == "" || password == "" {
		panic("username or password is required")
	}
	databaseName := os.Getenv("DB_NAME")
	ginMode := gin.Mode()
	var mongoDBUrl string
	if ginMode == "debug" {
		mongoDBUrl = "mongodb+srv://"+username+":"+password+"@cluster0.qq6u4.mongodb.net/"+databaseName+"?retryWrites=true&w=majority"
	}

	clientOptions := options.Client().ApplyURI(mongoDBUrl)

	client, dbErr := mongo.Connect(ctx, clientOptions)
	if dbErr != nil {
		panic("Failed to connect to database " + dbErr.Error())
	}

	dbErr = client.Ping(ctx, nil)
	if dbErr != nil {
		panic("Failed to connect to database " + dbErr.Error())
	}

	if databaseName == "" {
		panic("Please provide the database name")
	}

	database := client.Database(databaseName)

	return database
	
}