package task

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetTasksHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Retrieve tasks"})
}

func CreateTaskHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Task created"})
}

func UpdateTaskHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func DeleteTaskHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}