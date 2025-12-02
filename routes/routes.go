package routes

import (
	"github.com/gin-gonic/gin"
	"taskflow/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/tasks", controllers.CreateTask)
}