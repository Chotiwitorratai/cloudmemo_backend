package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // Adds some metadata fields to the table
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;one2many:user_id"` // Explicitly specify the type to be uuid
	Username   string `gorm:"uniqueIndex;not nul"`
	Email      string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	Image      *string
	Favorites  []Memo `gorm:"many2many:favorites;"`

}

type Memo struct {
	gorm.Model
	Slug        string `gorm:"uniqueIndex;not null;primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Body        string
	Author      User 
	AuthorID    uint
	Weather     string
	MusicUrl    *string
	Comments    []Comment
	Favorites   []User `gorm:"many2many:favorites;"`
}

type Comment struct {
	gorm.Model
	Memo   Memo
	ArticleID uint
	User      User
	UserID    uint
	Body      string
}