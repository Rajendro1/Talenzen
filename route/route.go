package route

import (
	"github.com/Rajendro1/Talenzen/handler/auth"
	"github.com/Rajendro1/Talenzen/handler/task"
	"github.com/Rajendro1/Talenzen/middleware"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	router := gin.Default()

	router.POST("/login", auth.LoginHandler)
	router.POST("/register", auth.RegisterHandler)
	router.GET("/tasks", middleware.AuthMiddleware(), task.GetTasksHandler)
	router.GET("/task_by_id/:task_id", middleware.AuthMiddleware(), task.GetTasksByIdHandler)
	router.GET("/task_by_users/:userID", middleware.AuthMiddleware(), task.GetTasksByUserIDHandler)
	router.POST("/tasks", middleware.AuthMiddleware(), task.CreateTaskHandler)
	router.POST("/assign_task", middleware.AuthMiddleware(), task.AssignTaskHandler)
	router.PUT("/tasks/", middleware.AuthMiddleware(), task.UpdateTaskHandler)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(), task.DeleteTaskHandler)

	router.Run(":8080")

}
