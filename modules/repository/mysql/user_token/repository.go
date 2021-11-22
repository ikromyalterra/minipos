package usertoken

import (
	"github.com/ikromyalterra/minipos/business/model"
	"github.com/jinzhu/gorm"
)

type (
	Repository struct {
		*gorm.DB
	}
)

func New(db *gorm.DB) *Repository {
	return &Repository{
		db.Table("user_token"),
	}
}

func (repo *Repository) Insert(userToken *model.UserToken) error {
	return repo.Create(userToken).Error
}

func (repo *Repository) GetByTokenID(tokenID uint) (model.UserToken, error) {
	var userToken model.UserToken
	result := repo.Where("id_token = ?", tokenID).Find(&userToken)
	return userToken, result.Error
}

func (db *Repository) DeleteByUserID(userID uint) error {
	return db.Where("id_user = ?", userID).Delete(&model.UserToken{}).Error
}

func (db *Repository) DeleteByTokenID(tokenID uint) error {
	return db.Where("id_token = ?", tokenID).Delete(&model.UserToken{}).Error
}
