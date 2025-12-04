package repositories

import (
	"taskflow/config"
	"taskflow/models"
)

type taskRepo struct{}

var TaskRepo = taskRepo{}

func (r taskRepo) Create(task *models.Task) error {
	return config.DB.Create(task).Error
}

func( r taskRepo) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	result := config.DB.Find(&tasks)
	return tasks, result.Error
}

func ( r taskRepo) FindByID(id uint) (models.Task, error) {
	var task models.Task
	result := config.DB.Find(&task, id)
	return task, result.Error
}

func( r taskRepo) Update(task *models.Task) error {
	result := config.DB.Save(task)
	return result.Error
}

func ( r taskRepo) Delete(id uint) error {
	result :=config.DB.Delete(&models.Task{}, id)
	return result.Error
}