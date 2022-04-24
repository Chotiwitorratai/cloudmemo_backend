package router

import (
	controllers "github.com/Chotiwitorratai/cloudmemo_backend/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Create routes group.
	user := app.Group("/user")
	memo := app.Group("/memo")

	// Routes for POST method:
	user.Post("/sign/up",  controllers.UserSignUp)           // create a new book
	user.Post("/sign/in",  controllers.UserSignIn)           // create a new book
	user.Get("/get/all",  controllers.GetAllUsers)           // create a new book
	// route.Post("/user/sign/out", middleware.JWTProtected(), controllers.UserSignOut) // de-authorization user
	// route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)   // renew Access & Refresh tokens
	memo.Post("/create", controllers.CreateMemo)
	memo.Get("/get/all/:user_id", controllers.GetAllMemo)
	memo.Get("/get/:memo_id", controllers.GetMemo)
	memo.Get("/get/shared/:memo_id", controllers.GetSharedToken)
	// // Routes for PUT method:
	// route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook) // update one book by ID

	// // Routes for DELETE method:
	// route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID
}