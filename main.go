package main

import (
	"log"

	"github.com/Chotiwitorratai/cloudmemo_backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRouter(app *fiber.Ctx){

}

func main() {
    app := fiber.New()
    database.ConnectDB()

    app.Use(logger.New(logger.Config{
        Format:"[${ip}]:${port} ${status} - ${method} ${path}\n",
    }))
    app.Get("/", func(c *fiber.Ctx) error {
    // send text
    return c.SendString("Hello, World!")
    })    

    log.Fatal(app.Listen(":3001"))
    // app.Listen(":3000")
}