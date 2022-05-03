package controllers

import (
	"fmt"

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
			"message":   err.Error(),
		})
	}
	claims, err := utils.ExtractTokenMetadata(c)
	CheckToken(c)
	db := database.DB
	db.Find(&users, "id = ?" ,claims.UserID )
	 if users.ID == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
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
    err = db.Create(&memo).Error
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Could not create note",
			"data": err,	
		})
    }
    // Return the created note
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"message": "Created Memo",
		"data": memo,
	})
	
}

func UpdateMemo(c *fiber.Ctx) error {
	um := &model.UpdateMemo{}
	err := c.BodyParser(um)
	if  err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message":   err.Error(),
		})
	}
	CheckToken(c)
	db := database.DB
	memo := &model.Memo{}
	db.First(&memo, um.ID)	
    memo.Title = um.Title
    memo.Description = um.Description
    memo.Body = um.Body
    memo.Weather = um.Weather
    memo.MusicUrl = um.MusicUrl
    err = db.Save(&memo).Error
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Could not create note",
			"data": err,	
		})
    }
	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Updated Memo",
		"data": memo,
	})
}

func PublishMemo(c *fiber.Ctx) error {
	um := &model.PublishMemo{}
	err := c.BodyParser(um)
	if  err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message":   err.Error(),
		})
	}
	CheckToken(c)
	db := database.DB
	memo := &model.Memo{}
	db.First(&memo, um.ID)	
    memo.IsPublic = !memo.IsPublic
    err = db.Save(&memo).Error
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Could not create note",
			"data": err,	
		})
    }
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Updated Memo",
		"data": memo,
	})
}

func DeleteMemo(c *fiber.Ctx) error {
	id := c.Params("memo_id")
	err  := CheckToken(c)
	db := database.DB
	memo := &model.Memo{}
    db.First(&memo ,id )
    
    err = db.Delete(&memo).Error
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Could not delete memo",
			"data": err,	
		})
    }
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Deleted memo",
		"data": memo,
	})
}

func GetAllMemo(c *fiber.Ctx) error {
	db := database.DB
	CheckToken(c)
    var memos []model.Memo
	fmt.Println(memos)
	id := c.Params("user_id")
    db.Find(&memos, "author_id = ?" ,id )
    if len(memos) == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": memos})
	
}

func GetMemo(c *fiber.Ctx) error {
	db := database.DB
    var memos []model.Memo
	id := c.Params("memo_id")
    db.First(&memos ,id )
    // If no user is present return an error
    if len(memos) == 0 {
		token,err := utils.ExtractToken(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
		}
		db.Find(&memos, "id = ?" ,token.UserID )
		if len(memos) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": memos})
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": memos})
}

func GetSharedToken(c *fiber.Ctx) error {
    var memos []model.Memo
	fmt.Println(memos)
	id := c.Params("memo_id")
	token, err := utils.GenerateNewSharedTokens(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}
    return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Found", "data": token})

}