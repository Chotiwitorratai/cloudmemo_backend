package controllers

import (
	"github.com/Chotiwitorratai/cloudmemo_backend/database"
	"github.com/Chotiwitorratai/cloudmemo_backend/model"
	"github.com/gofiber/fiber/v2"
)


func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
    var users []model.User

    db.Find(&users)

    // If no user is present return an error
    if len(users) == 0 {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
    }

    return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": users})
	
}

// func userIDFromToken(c *fiber.Ctx) uint {
// 	var user *jwt.Token
// 	l := c.Locals("user")
// 	if l == nil {
// 		return 0
// 	}
// 	user = l.(*jwt.Token)
// 	id := uint(((user.Claims.(jwt.MapClaims)["id"]).(float64)))
// 	return id
// }