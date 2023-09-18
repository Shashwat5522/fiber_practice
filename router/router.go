package router

import (
	"fiber/handlers"
	"fiber/middleware"

	

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.AuthReq())

	api.Get("/", handlers.GetAllProducts)
	api.Get("/:id", handlers.GetSingleProduct)
	api.Post("/", handlers.CreateProduct)
	api.Delete("/:id", handlers.DeleteProduct)
	api.Put("/:id",handlers.UpdateProduct)
}
