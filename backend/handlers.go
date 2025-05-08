package main

import (
    "context"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func getTasks(c *gin.Context) {
    var tasks []Task
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Printf("Error finding tasks: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
        return
    }
    defer cursor.Close(context.TODO())

    if err = cursor.All(context.TODO(), &tasks); err != nil {
        log.Printf("Error decoding tasks: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process tasks"})
        return
    }

    if len(tasks) == 0 {
        c.JSON(http.StatusOK, []Task{})
        return
    }

    c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
    var task Task
    if err := c.ShouldBindJSON(&task); err != nil {
        log.Printf("Invalid task JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
        return
    }

    if task.Title == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
        return
    }

    task.ID = primitive.NewObjectID().Hex()
    result, err := collection.InsertOne(context.TODO(), task)
    if err != nil {
        log.Printf("Error inserting task: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
        return
    }

    log.Printf("Created task with ID: %s", result.InsertedID)
    c.JSON(http.StatusCreated, task)
}

func updateTask(c *gin.Context) {
    id := c.Param("id")
    if !primitive.IsValidObjectID(id) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    var task Task
    if err := c.ShouldBindJSON(&task); err != nil {
        log.Printf("Invalid update JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update data"})
        return
    }

    if task.Status == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
        return
    }

    result, err := collection.UpdateOne(
        context.TODO(),
        bson.M{"_id": id},
        bson.M{"$set": bson.M{"status": task.Status}},
    )
    if err != nil {
        log.Printf("Error updating task %s: %v", id, err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
        return
    }

    if result.MatchedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    log.Printf("Updated task with ID: %s", id)
    c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func deleteTask(c *gin.Context) {
    id := c.Param("id")
    if !primitive.IsValidObjectID(id) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
    if err != nil {
        log.Printf("Error deleting task %s: %v", id, err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
        return
    }

    if result.DeletedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    log.Printf("Deleted task with ID: %s", id)
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}