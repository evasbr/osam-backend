package handler

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/evasbr/osam-backend/app/config"
	"github.com/evasbr/osam-backend/app/dto"
	"github.com/evasbr/osam-backend/app/errors"
	"github.com/evasbr/osam-backend/app/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var (
	once     sync.Once
	appFiber *fiber.App
)

func initApp() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.ConnectToDB()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			switch e := err.(type) {
			case errors.ValidationError:
				return c.Status(fiber.StatusBadRequest).JSON(dto.GlobalErrorHandlerResp{
					Error: e.Messages,
				})

			case errors.HttpError:
				return c.Status(e.StatusCode).JSON(dto.GlobalErrorHandlerResp{
					Error: e.Messages,
				})

			default:
				return c.Status(fiber.StatusInternalServerError).JSON(dto.GlobalErrorHandlerResp{
					Error: []string{"Internal Server Error"},
				})
			}
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	route.AuthRoutes(api)

	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(dto.GlobalErrorHandlerResp{
			Error: []string{"Route not found"},
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	appFiber = app
}

// Handler untuk Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(initApp) // Pastikan hanya inisialisasi satu kali
	adaptor.FiberApp(appFiber)(w, r)
}
