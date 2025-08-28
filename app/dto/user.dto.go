package dto

import "time"

type RegisterUserDTO struct {
	Name            string `validate:"required,min=3"`
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=6"`
	BirthDate       *time.Time
	TelephoneNumber *string
}

type LoginUserDTO struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
