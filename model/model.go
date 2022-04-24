package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // Adds some metadata fields to the table
	Username   string `gorm:"uniqueIndex;not nul"`
	Email      string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	Bio        *string
	Image      *string
	Favorites  []Memo `gorm:"many2many:favorites;"`

}


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
		// Slug        string `gorm:"uniqueIndex;not null"`
		Title       string `gorm:"not null"`
		Description string
		Body        string
		Weather     string
		MusicUrl    *string
		IsPublic    bool `gorm:"default:false"`
		Author      User
		AuthorID    uint
		Comments    []Comment
		Favorites   []User `gorm:"many2many:favorites;"`
		// Tags        []Tag  `gorm:"many2many:article_tags;"`
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

	type Comment struct {
	gorm.Model
	Memo      Memo
	MemoID    uint
	User      User
	UserID    uint
	Body      string
	}
