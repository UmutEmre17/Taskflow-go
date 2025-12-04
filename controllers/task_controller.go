package controllers

import (
	"net/http"
	"taskflow/services"
	"github.com/gin-gonic/gin"
	"strconv"
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

func GetTaskByID(c *gin.Context) {
	idParam := c.Param("id")

	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(id64)

	task, err := services.TaskService.GetTaskByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

type UpdateTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")

	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(id64)

	var body UpdateTaskRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}


	updatedTask, err := services.TaskService.UpdateTask(id, body.Title, body.Status)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")

	// id string → uint dönüşümü
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(id64)

	// service’i çağır
	err = services.TaskService.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}