package task

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Rajendro1/Talenzen/db/pgd"
	"github.com/Rajendro1/Talenzen/model"
	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	var newTask model.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pgd.CreateTask(newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create task"})
		return
	}
	c.Status(http.StatusCreated)
}

func GetTasksHandler(c *gin.Context) {
	tasks, err := pgd.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
func GetTasksByIdHandler(c *gin.Context) {
	taskId := c.Param("task_id")
	task, err := pgd.GetTasksByID(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve task", "details": "Please try again later."})
		log.Printf("Error retrieving task: %v", err)
		return
	}
	c.JSON(http.StatusOK, task)
}

func GetTasksByUserIDHandler(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tasks, err := pgd.GetTasksByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve tasks", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
func UpdateTaskHandler(c *gin.Context) {
	var t model.Task
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pgd.UpdateTask(t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update task"})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteTaskHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	if err := pgd.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete task"})
		return
	}
	c.Status(http.StatusOK)
}

func AssignTaskHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.PostForm("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	taskID, err := strconv.Atoi(c.PostForm("task_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if !pgd.UserExists(userID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !pgd.TaskExists(taskID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	var cl model.Claims
	if err := pgd.AssignTask(userID, taskID, cl.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign task", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task assigned successfully"})
}
