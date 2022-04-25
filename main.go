package main

import (
	"log"
	"os"

	"github.com/Chotiwitorratai/cloudmemo_backend/database"
	middleware "github.com/Chotiwitorratai/cloudmemo_backend/middleware"
	router "github.com/Chotiwitorratai/cloudmemo_backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)



func main() {
    app := fiber.New()
    database.ConnectDB()
    app.Use(logger.New(logger.Config{
        Format:"[${ip}]:${port} ${status} - ${method} ${path}\n",
    }))
    middleware.FiberMiddleware(app)
    app.Get("/", func(c *fiber.Ctx) error {
    // send text
    return c.SendString("Online! Send your API")
    })    
    app.Use(limiter.New())
	router.SetupRoutes(app)
    router.NotFoundRoute(app)
    // log.Fatal(app.Listen(":3000"))
    log.Fatal(app.Listen(":" + os.Getenv("PORT")))
    // app.Listen(":3000")
}