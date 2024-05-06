package repository

import (
	"app/internal/models"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	if r.DB == nil {
		log.Println("DATABASE IS NIL")
		return nil, nil
	}
	return nil, nil
}

// func (r *UserRepository) CreateUser(user *models.User) error {
// 	// Реализация метода создания пользователя в базе данных
// }

// func (r *UserRepository) UpdateUser(user *models.User) error {
// 	// Реализация метода обновления пользователя в базе данных
// }

// func (r *UserRepository) DeleteUser(id int) error {
// 	// Реализация метода удаления пользователя из базы данных
// }
