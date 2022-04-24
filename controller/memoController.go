package controllers

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Chotiwitorratai/cloudmemo_backend/database"
	"github.com/Chotiwitorratai/cloudmemo_backend/model"
	"github.com/Chotiwitorratai/cloudmemo_backend/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateMemo(c *fiber.Ctx) error {
	cm := &model.CreateMemo{}
	users := &model.User{}
	err := c.BodyParser(cm)
	if  err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"msg":   err.Error(),
		})
	}
	now := time.Now().Unix()
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"msg":   "unauthorized, check expiration time of your token",
		})
	}
	db := database.DB
	db.Find(&users, "id = ?" ,claims.UserID )
	 if users.ID == 0 {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
    }
    // var memo []model.Memo
	memo := &model.Memo{}
    // Add a uuid to the note
    memo.Title = cm.Title
    memo.Description = cm.Description
    memo.Body = cm.Body
    memo.Weather = cm.Weather
    memo.MusicUrl = cm.MusicUrl
    memo.AuthorID = claims.UserID
    // memo.IsPublic = false
    memo.Author = *users
    // Create the Note and return error if encountered
    err = db.Create(&memo).Error

    if err != nil {
        return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Could not create note",
			"data": err,	
		})
    }
    // Return the created note
    return c.JSON(fiber.Map{
		"status": "success",
		"message": "Created Note",
		"data": memo,
	})
	
}

func GetAllMemo(c *fiber.Ctx) error {
	db := database.DB
    var memos []model.Memo
	fmt.Println(memos)
	id := c.Params("user_id")
    db.Find(&memos, "author_id = ?" ,id )
    // If no user is present return an error
    if len(memos) == 0 {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
    }

    return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": memos})
	
}

func GetMemo(c *fiber.Ctx) error {
	db := database.DB
    var memos []model.Memo
	fmt.Println(memos)
	id := c.Params("memo_id")
	fmt.Println("var1 = ", reflect.TypeOf(id))
    db.Find(&memos, "id = ?" ,id )
    // If no user is present return an error
    if len(memos) == 0 {
		token,err := utils.ExtractToken(id)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
		}
		db.Find(&memos, "id = ?" ,token.UserID )
		if len(memos) == 0 {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
		}
		return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": memos})
    }
    return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": memos})
}

func GetSharedToken(c *fiber.Ctx) error {
    var memos []model.Memo
	fmt.Println(memos)
	id := c.Params("memo_id")
	token, err := utils.GenerateNewSharedTokens(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}
    return c.JSON(fiber.Map{"status": "success", "message": "Found", "data": token})

}