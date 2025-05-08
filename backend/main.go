package main

import (
    "context"
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
    ID          string `json:"_id" bson:"_id"`
    Title       string `json:"title" bson:"title"`
    Description string `json:"description" bson:"description"`
    Status      string `json:"status" bson:"status"`
}

var collection *mongo.Collection

func main() {
    // Load .env file if it exists, but don't crash if it doesn't
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on environment variables")
    }

    // Get and validate MONGO_URI
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI not set")
    }
    log.Println("Connecting to MongoDB at:", mongoURI)

    // Connect to MongoDB
    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }
    // Ensure MongoDB client is disconnected on exit
    defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            log.Printf("Error disconnecting MongoDB client: %v", err)
        }
    }()

    // Verify MongoDB connection
    if err := client.Ping(context.TODO(), nil); err != nil {
        log.Fatal("Failed to ping MongoDB:", err)
    }

    // Set up MongoDB collection
    collection = client.Database("taskdb").Collection("tasks")

    // Initialize Gin router
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

    // Define API routes
    r.GET("/api/tasks", getTasks)
    r.POST("/api/tasks", createTask)
    r.PUT("/api/tasks/:id", updateTask)
    r.DELETE("/api/tasks/:id", deleteTask)

    // Start server
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}