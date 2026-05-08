package usecase

import (
	"be-interview-app/internal/entity"
	"be-interview-app/internal/repository"
	"errors"

	customError "be-interview-app/internal/delivery/error"
	"be-interview-app/internal/delivery/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase interface {
	Register(name, password string) (*entity.User, error)
	UserProfile(id int) (*entity.User, error)
	Login(name, password string) (string, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Register(name, password string) (*entity.User, error) {

	existingUser, err := u.repo.FindByUsername(&name);

	if err == nil && existingUser.ID != 0 {
		return nil, customError.ErrDuplicateUser
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	
	// 1. เข้ารหัส Password (bcrypt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 2. สร้าง Entity User
	user := &entity.User{
		Name:     name,
		Password: string(hashedPassword),
	}

	// 3. สั่ง Repository ให้บันทึก
	if err := u.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}


func (u *userUsecase) UserProfile(id int) (*entity.User, error) {
	user, err := u.repo.FindById(&id)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Login(name, password string) (string, error) {

	user, err := u.repo.FindByUsername(&name)
	if err != nil {
		return "", customError.ErrInvalidCredential
	}

	// compare password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", customError.ErrInvalidCredential
	}

	// generate jwt
	token, err := utils.GenerateJWT(user.ID, user.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}