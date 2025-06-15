package auth

import (
	"github.com/bintangyosua/echaaas/internal/user"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Create(user *user.User) error {
	return r.DB.Create(user).Error
}

func (r *Repository) FindAll() ([]user.User, error) {
	var users []user.User
	err := r.DB.Find(&users).Error
	return users, err
}