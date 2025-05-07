package main

import (
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "net/http"
)

func getTasks(c *gin.Context) {
    var tasks []Task
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if err = cursor.All(context.TODO(), &tasks); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
    var task Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.ID = primitive.NewObjectID().Hex()
    _, err := collection.InsertOne(context.TODO(), task)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, task)
}

func updateTask(c *gin.Context) {
    id := c.Param("id")
    var task Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    _, err := collection.UpdateOne(
        context.TODO(),
        bson.M{"_id": id},
        bson.M{"$set": bson.M{"status": task.Status}},
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func deleteTask(c *gin.Context) {
    id := c.Param("id")
    _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}