package repository

import (
	"todo-app/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo models.Todo) (models.Todo, error)
	DeleteTodo(Todo models.Todo) (models.Todo, error)
	GetTodo(ID int) (models.Todo, error)
}

func repositoryTodo(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetTodo(ID int) (models.Todo, error) {
	var todo models.Todo
	err := r.db.Raw("SELECT * FROM products WHERE id=?", ID).Scan(&todo).Error
	return todo, err
}

func (r *repository) CreateTodo(Todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&Todo).Error
	return Todo, err
}

func (r *repository) DeleteTodo(Todo models.Todo) (models.Todo, error) {
	err := r.db.Delete(&Todo).Error

	return Todo, err
}
