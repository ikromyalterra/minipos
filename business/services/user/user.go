package user

import (
	"errors"

	"github.com/ikromyalterra/minipos/business/model"
	portMerchant "github.com/ikromyalterra/minipos/business/port/merchant"
	portOutlet "github.com/ikromyalterra/minipos/business/port/outlet"
	port "github.com/ikromyalterra/minipos/business/port/user"
	portUserToken "github.com/ikromyalterra/minipos/business/port/user_token"
	"github.com/ikromyalterra/minipos/utils/crypto"
)

type (
	service struct {
		userRepo      port.Repository
		merchantRepo  portMerchant.Repository
		outletRepo    portOutlet.Repository
		userTokenRepo portUserToken.Repository
	}
)

func New(userRepo port.Repository,
	merchantRepo portMerchant.Repository,
	outletRepo portOutlet.Repository,
	userTokenRepo portUserToken.Repository) port.Service {
	return &service{
		userRepo,
		merchantRepo,
		outletRepo,
		userTokenRepo,
	}
}

func (s *service) Create(user *port.User) error {
	if err := s.bindUserRole(user); err != nil {
		return err
	}

	data := new(model.User)
	data.Email = user.Email
	data.Role = user.Role
	data.MerchantID = user.MerchantID
	data.OutletID = user.OutletID
	data.Password = crypto.UserGeneratePassword(user.Password)

	if err := s.userRepo.Insert(data); err != nil {
		return err
	}
	user.ID = data.ID

	return nil
}

func (s *service) GetByID(ID uint) (port.User, error) {
	var data port.User

	user, err := s.userRepo.GetByID(ID)
	if data.ID > 0 {
		data.ID = user.ID
		data.Email = user.Email
		data.Password = user.Password
		data.MerchantID = user.MerchantID
		data.OutletID = user.OutletID
	}
	return data, err
}

func (s *service) View(ID uint) (port.User, error) {
	var data port.User

	user, err := s.userRepo.View(ID)
	if user.ID > 0 {
		data.ID = user.ID
		data.Email = user.Email
		data.Role = user.Role
		data.MerchantID = user.MerchantID
		data.MerchantName = user.MerchantName
		data.OutletID = user.OutletID
		data.OutletName = user.OutletName
	}

	return data, err
}

func (s *service) Update(user *port.User) error {
	existingUser, err := s.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}
	if existingUser.ID == 0 {
		return errors.New("id not found")
	}
	if err := s.bindUserRole(user); err != nil {
		return err
	}

	var data model.User
	data.ID = user.ID
	data.Email = user.Email
	data.Role = user.Role
	data.MerchantID = user.MerchantID
	data.Password = existingUser.Password
	data.OutletID = user.OutletID
	data.CreatedAt = existingUser.CreatedAt

	revokeToken := false

	if user.Password != "" {
		data.Password = crypto.UserGeneratePassword(user.Password)
		revokeToken = true
	}
	if !revokeToken {
		revokeToken = existingUser.MerchantID != user.MerchantID || existingUser.OutletID != user.OutletID
	}

	err = s.userRepo.Update(data)
	if err == nil && revokeToken {
		s.userTokenRepo.DeleteByUserID(user.ID)
	}

	return err
}

func (s *service) Delete(ID uint) error {
	return s.userRepo.Delete(ID)
}

func (s *service) List() ([]port.User, error) {
	datas := make([]port.User, 0)

	users, err := s.userRepo.List()
	if err != nil {
		return datas, err
	}

	var user port.User
	for i := range users {
		user.ID = users[i].ID
		user.Email = users[i].Email
		user.Role = users[i].Role
		user.MerchantID = users[i].MerchantID
		user.MerchantName = users[i].MerchantName
		user.OutletID = users[i].OutletID
		user.OutletName = users[i].OutletName

		datas = append(datas, user)
	}

	return datas, nil
}

func (s *service) bindUserRole(user *port.User) error {
	switch user.Role {
	case "merchant":
		user.OutletID = 0
		currentMerchant, err := s.merchantRepo.GetByID(user.MerchantID)
		if err != nil || currentMerchant.ID == 0 {
			return errors.New("id_merchant not found")
		}
	case "outlet":
		user.MerchantID = 0
		currentOutlet, err := s.outletRepo.GetByID(user.OutletID)
		if err != nil || currentOutlet.ID == 0 {
			return errors.New("id_outlet not found")
		}
	default:
		user.MerchantID = 0
		user.OutletID = 0
	}
	return nil
}
