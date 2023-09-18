package middleware

import (
	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
)

func AuthReq() func(*fiber.Ctx) {
	cfg := basicauth.Config{
		Users: map[string]string{
			"root": "root",
		},
	}
	err := basicauth.New(cfg)
	return err

}
