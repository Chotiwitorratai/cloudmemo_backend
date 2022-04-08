package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    gorm.Model           
    user_id         uuid.UUID `gorm:"type:uuid"` 
    name      string
    surname   string
}

type Memo struct {
	gorm.Model           
    memo_id         uuid.UUID `gorm:"type:uuid"`
    detail      string
}