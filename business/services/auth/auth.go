package auth

import (
	"errors"

	"github.com/ikromyalterra/minipos/business/model"

	port "github.com/ikromyalterra/minipos/business/port/auth"
	userPort "github.com/ikromyalterra/minipos/business/port/user"
	userTokenPort "github.com/ikromyalterra/minipos/business/port/user_token"
	"github.com/ikromyalterra/minipos/utils/auth"
	"github.com/ikromyalterra/minipos/utils/crypto"
	"github.com/ikromyalterra/minipos/utils/helper"
)

type (
	service struct {
		userTokenRepository userTokenPort.Repository
		userRepository      userPort.Repository
	}
)

func New(userTokenRepository userTokenPort.Repository, userRepository userPort.Repository) port.Service {
	return &service{
		userTokenRepository,
		userRepository,
	}
}

var (
	jwt                  auth.JWT = auth.NewJWT()
	ErrGenerateToken     error    = errors.New("generate token failed")
	ErrInvalidToken      error    = errors.New("invalid token")
	ErrInvalidCredential error    = errors.New("invalid credential")
)

func (s *service) Verify(tokenString string) (interface{}, error) {
	token, tokenClaims, err := jwt.Parse(tokenString)
	if err != nil {
		return nil, err
	}
	userToken, err := s.userTokenRepository.GetByTokenID(tokenClaims.ID)
	if userToken.UserID == 0 || err != nil {
		return nil, ErrInvalidToken
	}
	return token, nil
}

func (s *service) Login(authUser *port.AuthUser) error {
	if !s.bindUserCredential(authUser) {
		return ErrInvalidCredential
	}

	tokenID := helper.GenerateRandomNumber()
	if err := generateToken(tokenID, authUser); err == nil {
		userToken := new(model.UserToken)
		userToken.TokenID = tokenID
		userToken.UserID = authUser.UserID

		return s.userTokenRepository.Insert(userToken)
	}

	return ErrGenerateToken
}

func (s *service) Logout(tokenID uint) error {
	return s.userTokenRepository.DeleteByTokenID(tokenID)
}

func generateToken(tokenID uint, authUser *port.AuthUser) (err error) {
	var tokenClaims auth.JWTClaims
	tokenClaims.ID = tokenID
	tokenClaims.Email = authUser.Email
	tokenClaims.Role = authUser.Role

	authUser.Token, err = jwt.Create(tokenClaims)

	return
}

func (s *service) bindUserCredential(authUser *port.AuthUser) bool {
	existingUser, _ := s.userRepository.GetByEmail(authUser.Email)
	if existingUser.ID > 0 {
		authUser.UserID = existingUser.ID
		authUser.Role = existingUser.Role
		return crypto.UserVerifyPassword(authUser.Password, existingUser.Password)
	}

	return false
}
