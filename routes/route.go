package router

import (
	controllers "github.com/Chotiwitorratai/cloudmemo_backend/controller"
	"github.com/Chotiwitorratai/cloudmemo_backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Create routes group.
	user := app.Group("/user")
	memo := app.Group("/memo")
	token := app.Group("/token")
	youtube := app.Group("/youtube")

	user.Post("/sign/up",  controllers.UserSignUp)           
	user.Post("/sign/in",  controllers.UserSignIn)
	user.Get("/get/all",  controllers.GetAllUsers)
	user.Get("/get/:user_id",  controllers.GetUser)
	// route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)   // renew Access & Refresh tokens
	memo.Post("/create",middleware.JWTProtected(), controllers.CreateMemo)
	memo.Put("/update",middleware.JWTProtected(), controllers.UpdateMemo)
	memo.Put("/publish",middleware.JWTProtected(), controllers.PublishMemo)
	memo.Delete("/delete/:memo_id",middleware.JWTProtected(), controllers.DeleteMemo)
	memo.Get("/get/all/:user_id",middleware.JWTProtected(), controllers.GetAllMemo)
	memo.Get("/get/:memo_id", controllers.GetMemo)
	memo.Get("/get/shared/:memo_id",middleware.JWTProtected(), controllers.GetSharedToken)

	token.Post("/renew", middleware.JWTProtected(), controllers.RenewTokens)
	youtube.Post("/Search", middleware.JWTProtected(), controllers.SearchMusic)
}