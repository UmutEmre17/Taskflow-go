package controllers

import (
	"net/http"
	"taskflow/services"
	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Title string `json:"title"`
}

func CreateTask(c *gin.Context) {
	var body CreateTaskRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	task, err := services.TaskService.CreateTask(body.Title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "task not created"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func GetTasks(c *gin.Context) {
	tasks, err := services.TaskService.GetAllTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}