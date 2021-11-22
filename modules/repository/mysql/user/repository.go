package user

import (
	"errors"
	"reflect"

	"github.com/ikromyalterra/minipos/business/model"
	"github.com/ikromyalterra/minipos/utils/helper"
	"github.com/jinzhu/gorm"
)

type (
	Repository struct {
		*gorm.DB
	}
)

func New(db *gorm.DB) *Repository {
	return &Repository{
		db.Table("user"),
	}
}

func (repo *Repository) Insert(user *model.User) error {
	return repo.Create(user).Error
}

func (repo *Repository) GetByID(ID uint) (model.User, error) {
	var user model.User
	result := repo.Where("id = ?", ID).Find(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return user, result.Error
}

func (repo *Repository) GetByEmail(email string) (model.User, error) {
	var user model.User
	result := repo.Where("email = ?", email).Find(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return user, result.Error
}

func (repo *Repository) View(ID uint) (model.UserView, error) {
	var user model.UserView
	result := repo.Select("user.id, user.email, user.password, user.role, IFNULL(user.id_merchant, mo.id) as id_merchant, user.id_outlet, user.created_at, user.updated_at, IFNULL(merchant.name, mo.name) as merchant, outlet.name as outlet").
		Joins("left join merchant on user.id_merchant = merchant.id").
		Joins("left join outlet on user.id_outlet = outlet.id").
		Joins("left join merchant as mo on outlet.id_merchant = mo.id").
		Where("user.id = ?", ID).
		Find(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return user, result.Error
}

func (repo *Repository) Update(user model.User) error {
	mapUser, err := helper.StructToMap(user)
	if err != nil {
		return err
	}
	if reflect.ValueOf(mapUser["id_merchant"]).IsZero() {
		mapUser["id_merchant"] = gorm.Expr("NULL")
	}
	if reflect.ValueOf(mapUser["id_outlet"]).IsZero() {
		mapUser["id_outlet"] = gorm.Expr("NULL")
	}
	return repo.Model(&user).Updates(mapUser).Error
}

func (repo *Repository) Delete(ID uint) error {
	result := repo.Where("id = ?", ID).Delete(&model.User{})
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return result.Error
}

func (repo *Repository) List() ([]model.UserView, error) {
	var users []model.UserView
	repo.Select("user.id, user.email, user.password, user.role, IFNULL(user.id_merchant, mo.id) as id_merchant, user.id_outlet, user.created_at, user.updated_at, IFNULL(merchant.name, mo.name) as merchant, outlet.name as outlet").
		Joins("left join merchant on user.id_merchant = merchant.id").
		Joins("left join outlet on user.id_outlet = outlet.id").
		Joins("left join merchant as mo on outlet.id_merchant = mo.id").
		Find(&users)
	return users, nil
}
