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