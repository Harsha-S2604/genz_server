package routes

import (
	"genz_server/service/userservice"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
	
	router := gin.Default()
	router.Use(cors.Default())

	userAPIRouter := router.Group("api/v1/users")
	{
		userAPIRouter.POST("/register", userservice.AddUserHandler(db))

		userAPIRouter.POST("/login", userservice.UserLoginHandler(db))

		userAPIRouter.GET("/user/:id", userservice.GetUserByIdHandler(db))

	}

	return router

}