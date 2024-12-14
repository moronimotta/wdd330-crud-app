package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/moronimotta/wdd330-crud-app/http"
	"github.com/moronimotta/wdd330-crud-app/repository"
)

func main() {
	mongoURI := ""
	if os.Getenv("ENV") != "production" {

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		mongoURI = os.Getenv("MONGO_URI")
		if mongoURI == "" {
			log.Fatal("MONGO_URI not set in .env file")
		}
	} else {
		mongoURI = os.Getenv("MONGO_URI")
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

	// Enable CORS for all origins
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	{
		router.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Hello, World!"})
		})
		// login
		router.GET("/users/:email/:password", server.GetUser)

		router.GET("/users/:email", server.GetUserByEmail)
		// register
		router.POST("/users", server.CreateUser)
		router.PUT("/users/:email", server.UpdateUser)

		router.GET("/meal-plans/:id", server.GetMeal)
		router.POST("/meal-plans", server.CreateMeal)
		router.PUT("/meal-plans/:id", server.UpdateMeal)
	}

	// start the router
	router.Run(":9080")
}
