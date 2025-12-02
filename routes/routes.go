package routes

import (
	"github.com/gin-gonic/gin"
	"taskflow/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/create-task", controllers.CreateTask)
	r.GET("/tasks", controllers.GetTasks)
}