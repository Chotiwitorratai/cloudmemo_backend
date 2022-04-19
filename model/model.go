package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // Adds some metadata fields to the table
	UserID         uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	Username   string `gorm:"uniqueIndex;not nul"`
	Email      string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	Bio        *string
	Image      *string
	Favorites  []Memo `gorm:"many2many:favorites;"`

}

// type Memo struct {
// 	gorm.Model
// 	ID	uuid.UUID `gorm:"type:uuid"`
// 	Slug        string `gorm:"uniqueIndex;not null;"`
// 	Title       string `gorm:"not null"`
// 	Description string
// 	Body        string
// 	Author      User
// 	AuthorID    uint
// 	Weather     string
// 	MusicUrl    *string
// 	Comments    []Comment `gorm:"foreignKey:MEMOID"`
// 	Favorites   []User `gorm:"many2many:favorites;"`
// }

// type Comment struct {
// 	gorm.Model
// 	MEMOID    uint
// 	Memo      Memo
// 	UserID    uint
// 	User      User
// 	Body      string
// }

type Memo struct {
	gorm.Model
	Slug        string `gorm:"uniqueIndex;not null"`
	Title       string `gorm:"not null"`
	Description string
	Body        string
	Weather     string
	MusicUrl    *string
	Author      User
	AuthorID    uint
	Comments    []Comment
	Favorites   []User `gorm:"many2many:favorites;"`
	// Tags        []Tag  `gorm:"many2many:article_tags;"`
}


type Comment struct {
	gorm.Model
	Memo   Memo
	MemoID uint
	User      User
	UserID    uint
	Body      string
}

// type Tag struct {
// 	gorm.Model
// 	Tag      string    `gorm:"uniqueIndex"`
// 	Articles []Memo `gorm:"many2many:article_tags;"`
// }