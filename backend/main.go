package main

import (
    "context"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "os"
)

type Task struct {
    ID          string `json:"_id" bson:"_id"`
    Title       string `json:"title" bson:"title"`
    Description string `json:"description" bson:"description"`
    Status      string `json:"status" bson:"status"`
}

var collection *mongo.Collection

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGO_URI")
    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection = client.Database("taskdb").Collection("tasks")

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

    r.GET("/api/tasks", getTasks)
    r.POST("/api/tasks", createTask)
    r.PUT("/api/tasks/:id", updateTask)
    r.DELETE("/api/tasks/:id", deleteTask)

    r.Run(":8080")
}