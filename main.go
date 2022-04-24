package main

import (
	"log"

	"github.com/Chotiwitorratai/cloudmemo_backend/database"
	router "github.com/Chotiwitorratai/cloudmemo_backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)



func main() {
    app := fiber.New()
    d := database.ConnectDB()
    go database.AutoMigrate(d)
    app.Use(logger.New(logger.Config{
        Format:"[${ip}]:${port} ${status} - ${method} ${path}\n",
    }))
    app.Get("/", func(c *fiber.Ctx) error {
    // send text
    return c.SendString("Online! Send your API")
    })    

	router.SetupRoutes(app)
    log.Fatal(app.Listen(":3000"))
    // app.Listen(":3000")
}