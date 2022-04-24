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

func (mo *MemoDB) CreateMemo(m *model.Memo) (err error) {
	return mo.db.Create(m).Error
}

func (mo *MemoDB) UpdateMemo(m *model.Memo) error {
	return mo.db.Model(m).Updates(m).Error
}


