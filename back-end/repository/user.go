package repository

import (
	"todo-app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(UserID int) (models.User, error)
	Login(email string) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) CreateUser(User models.User) (models.User, error) {
	err := r.db.Create(&User).Error
	return User, err
}
func (r *repository) GetUser(UserID int) (models.User, error) {
	var User models.User
	err := r.db.Preload("Todo").First(&User, UserID).Error
	return User, err
}
