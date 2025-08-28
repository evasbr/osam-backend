package route

import (
	"github.com/evasbr/osam-backend/app/config"
	"github.com/evasbr/osam-backend/app/controller"
	"github.com/evasbr/osam-backend/app/service"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	authService := service.NewAuthService(config.DB)
	authController := controller.NewAuthController(*authService)

	api := router.Group("/auth")
	api.Post("/register", authController.Register)

	api.Post("/login", authController.Login)

	// router.Post("/logout", controller.Logout)
}
