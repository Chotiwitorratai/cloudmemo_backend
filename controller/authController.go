package controllers

import (
	"time"

	"github.com/Chotiwitorratai/cloudmemo_backend/database"
	"github.com/Chotiwitorratai/cloudmemo_backend/model"
	"github.com/Chotiwitorratai/cloudmemo_backend/utils"
	"github.com/gofiber/fiber/v2"
)


func UserSignUp(c *fiber.Ctx) error {
	signUp := &model.SignUp{}
	err := c.BodyParser(signUp)
	if err := c.BodyParser(signUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db := database.DB
	// Create a new user struct.
	user := &model.User{}

	user.CreatedAt = time.Now()
	user.Email = signUp.Email
	user.Username = signUp.Username
	user.Password = utils.GeneratePassword(signUp.Password)
	user.Image = signUp.Image

	err = db.Create(user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	// Return the created user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Created Note", "data": user})
}

func UserSignIn(c *fiber.Ctx) error {
	signIn := &model.SignIn{}
	user := &model.User{}
	err := c.BodyParser(signIn)
	if err := c.BodyParser(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.DB
	db.Find(&user,"email = ?",signIn.Email)
	if user.ID == 0 {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
    }
	compareUserPassword := utils.ComparePasswords(user.Password, signIn.Password)
	if !compareUserPassword {
		// Return, if password is not compare to stored in database.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"msg":   "wrong user email address or password",
		})
	}
	tokens, err := utils.GenerateNewTokens(user.ID)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"msg":   err.Error(),
		})
	}
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"status": "success",
		"msg":   nil,
		"tokens": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})
	
}