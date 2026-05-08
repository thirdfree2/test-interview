package repository

import (
	"be-interview-app/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByUsername(name *string) (entity.User, error)
	FindById(id *int) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	// GORM จะใช้ข้อมูลใน struct บันทึกลง Table "users"
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUsername(name *string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("name = ?", *name).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindById(id *int) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", *id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}