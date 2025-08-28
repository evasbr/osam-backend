package controller

import (
	"github.com/evasbr/osam-backend/app/dto"
	"github.com/evasbr/osam-backend/app/errors"
	"github.com/evasbr/osam-backend/app/service"
	"github.com/evasbr/osam-backend/app/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var input dto.RegisterUserDTO

	if err := ctx.BodyParser(&input); err != nil {
		return errors.HttpError{
			StatusCode: fiber.StatusBadRequest,
			Messages:   []string{"Invalid request body"},
		}
	}

	if err := utils.ValidateStruct(input); err != nil {
		return errors.ValidationError{
			Messages: err,
		}
	}

	user, token, err := c.Service.Register(input)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user,
		"token":   token,
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var input dto.LoginUserDTO

	if err := ctx.BodyParser(&input); err != nil {
		return errors.HttpError{
			StatusCode: 400,
			Messages:   []string{"Invalid request body"},
		}
	}

	if errs := utils.ValidateStruct(input); errs != nil {
		return errors.ValidationError{
			Messages: errs,
		}
	}

	user, token, err := c.Service.Login(input)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Login success",
		"user":    user,
		"token":   token,
	})
}

// func Logout(c *fiber.Ctx) error {
// 	// c.Cookie(&fiber.Cookie{
// 	// 	Name:     "token",
// 	// 	Value:    "",
// 	// 	Expires:  time.Now().Add(-time.Hour),
// 	// 	HTTPOnly: true,
// 	// 	Secure:   true,
// 	// 	SameSite: "None",
// 	// 	Path:     "/",
// 	// })

// 	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 	// 	"message": "Logged out successfully",
// 	// })
// }
