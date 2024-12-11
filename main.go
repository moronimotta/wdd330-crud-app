package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/moronimotta/wdd330-crud-app/http"
	"github.com/moronimotta/wdd330-crud-app/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	userRepo := repository.NewUserRepository(client.Database("wdd330-fn-project"))
	mealRepository := repository.NewMealPlanRepository(client.Database("wdd330-fn-project"))

	server := http.NewServer(userRepo, mealRepository)

	router := gin.Default()
	{
		router.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Hello, World!"})
		})

		router.GET("/users", server.ListUsers)
		router.GET("/users/:email", server.GetUser)
		router.POST("/users", server.CreateUser)
		router.PUT("/users/:email", server.UpdateUser)
		router.DELETE("/users/:email", server.DeleteUser)

		router.GET("/meal-plans", server.ListMeals)
		router.GET("/meal-plans/:id", server.GetMeal)
		router.POST("/meal-plans", server.CreateMeal)
		router.PUT("/meal-plans/:id", server.UpdateMeal)
		router.DELETE("/meal-plans/:id", server.DeleteMeal)
	}

	// start the router
	router.Run(":9080")
}
