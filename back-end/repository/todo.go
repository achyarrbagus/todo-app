package repository

import (
	"todo-app/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo models.Todo) (models.Todo, error)
}

func repositoryTodo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTodo(Todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&Todo).Error
	return Todo, err
}
