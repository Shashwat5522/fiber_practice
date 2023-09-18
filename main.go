package main

import (
	"fiber/database"
	"fiber/router"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	_ "github.com/lib/pq"
)

func main() {
	// fmt.Println("fiber")
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Use(middleware.Logger())
	router.SetupRoutes(app)
	app.Listen(3000)

}
