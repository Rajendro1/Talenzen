package route

import (
	"github.com/Rajendro1/Talenzen/handler/auth"
	"github.com/Rajendro1/Talenzen/handler/task"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	router := gin.Default()

	router.POST("/login", auth.LoginHandler)
	router.POST("/register", auth.RegisterHandler)
	router.GET("/tasks", task.GetTasksHandler)
	router.POST("/tasks", task.CreateTaskHandler)
	router.PUT("/tasks/:id", task.UpdateTaskHandler)
	router.DELETE("/tasks/:id", task.DeleteTaskHandler)

	router.Run(":8080")

}
