package authRoutes

import (
	authController "at3-back/internal/auth/infrastructure/controller"

	"github.com/gofiber/fiber/v2"
)

func Init(r fiber.Router) {
	var c authController.Auth
	c.New()

	r.Post("/login", c.Login)
	r.Post("/register", c.Register)
	r.Get("/confirm_account", c.Confirm_account)
	r.Post("/reset", c.Reset)
	r.Post("/register/company", c.RegisterCompany)
}
