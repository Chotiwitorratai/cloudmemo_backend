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
    if len(users) == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": users})
	
}

func GetUser(c *fiber.Ctx) error {
    CheckToken(c)
	db := database.DB
    id := c.Params("user_id")
    user := model.User{}
    db.First(&user,id)
    if user.ID == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": nil, "data": user})
	
}
