package services

import (
	"time"
	"taskflow/models"
	"taskflow/repositories"
)

type taskService struct {}

var TaskService = taskService{}

func (s taskService) CreateTask(title string) (models.Task, error) {
	task := models.Task{
		Title: title,
		Status: "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repositories.TaskRepo.Create(&task)
	return task, err
}

func (s taskService) GetAllTasks() ([]models.Task, error) {
	return repositories.TaskRepo.FindAll()
}

func (s taskService) GetTaskByID(id uint) (models.Task, error) {
	return repositories.TaskRepo.FindByID(id)
}