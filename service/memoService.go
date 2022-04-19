package service

import (
	"github.com/Chotiwitorratai/cloudmemo_backend/model"
	"gorm.io/gorm"
)

type MemoDB struct {
	db *gorm.DB
}

func NewMemoDB(db *gorm.DB) *MemoDB {
	return &MemoDB{
		db: db,
	}
}

func (mo *MemoDB) CreateMemoDB(a *model.Memo) error {
	tx := mo.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(a.ID).Preload("Favorites").Preload("Author").First(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}


func (us *UserDB) CreateMemo(u *model.Memo) (err error) {
	return us.db.Create(u).Error
}

func (us *UserDB) UpdateMemo(u *model.Memo) error {
	return us.db.Model(u).Updates(u).Error
}


