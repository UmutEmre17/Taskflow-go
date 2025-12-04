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

func (s taskService) UpdateTask(id uint, title string, status string) (models.Task, error) {
	task, err := repositories.TaskRepo.FindByID(id)

	if err != nil {
		return task, err
	}

	task.Title = title
	task.Status = status

	err = repositories.TaskRepo.Update(&task)
	return task, err
}

func(s taskService) DeleteTask(id uint) error {

	_, err := repositories.TaskRepo.FindByID(id)
	if err != nil {
		return err
	}

	return repositories.TaskRepo.Delete(id)
}