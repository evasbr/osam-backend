package service

import (
	"fmt"

	"github.com/evasbr/osam-backend/app/dto"
	"github.com/evasbr/osam-backend/app/errors"
	"github.com/evasbr/osam-backend/app/model"
	"github.com/evasbr/osam-backend/app/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) Register(input dto.RegisterUserDTO) (model.User, string, error) {
	var existingAuth model.User
	err := s.DB.Where("email = ?", input.Email).First(&existingAuth).Error

	if err == nil {
		return model.User{}, "", errors.HttpError{
			StatusCode: 400,
			Messages:   []string{"Email already used"},
		}
	}

	hashedPass, err := utils.HashPassword(input.Password)
	if err != nil {
		return model.User{}, "", err
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPass,
	}

	result := s.DB.Create(&user)
	if result.Error != nil {
		return model.User{}, "", errors.HttpError{
			StatusCode: 500,
			Messages:   []string{result.Error.Error()},
		}
	}

	token, err := utils.GenerateToken(user.ID.String())

	if err != nil {
		return model.User{}, "", errors.HttpError{
			StatusCode: 500,
			Messages:   []string{"Failed to generate token"},
		}
	}

	return user, token, nil
}

func (s *AuthService) Login(input dto.LoginUserDTO) (model.User, string, error) {
	var existingUser model.User
	err := s.DB.Where("email = ?", input.Email).First(&existingUser).Error
	if err != nil {
		fmt.Println("hee")
		return model.User{}, "", errors.HttpError{
			StatusCode: 400,
			Messages:   []string{"Email or password wrong"},
		}
	}

	if !utils.ComparePassword(input.Password, existingUser.Password) {
		return model.User{}, "", errors.HttpError{
			StatusCode: 400,
			Messages:   []string{"Email or password wrong"},
		}
	}

	token, err := utils.GenerateToken(existingUser.ID.String())

	if err != nil {
		return model.User{}, "", errors.HttpError{
			StatusCode: 500,
			Messages:   []string{"Failed to generate token"},
		}
	}

	return existingUser, token, nil
}
