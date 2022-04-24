package service

import (
	"errors"

	"github.com/Chotiwitorratai/cloudmemo_backend/model"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (us *UserDB) GetByID(id uint) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserDB) GetByEmail(e string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{Email: e}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserDB) GetByUsername(username string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{Username: username}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserDB) Create(u *model.User) (err error) {
	return us.db.Create(u).Error
}

func (us *UserDB) Update(u *model.User) error {
	return us.db.Model(u).Updates(u).Error
}


