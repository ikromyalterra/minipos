package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ikromyalterra/minipos/utils/config"
)

type (
	JWT interface {
		Create(tokenClaims JWTClaims) (string, error)
		Parse(tokenString string) (interface{}, *JWTClaims, error)
	}
	JWTConf   config.JWTConfig
	JWTClaims struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
		Role  string `json:"role"`
		jwt.StandardClaims
	}
)

var (
	ErrInvalidJWT error = errors.New("invalid or expired token")
)

func NewJWT() JWT {
	return &JWTConf{
		SignKey: config.GetJWTSigKey(),
		Expired: config.GetJWTExpired(),
	}
}

func (j *JWTConf) Create(tokenClaims JWTClaims) (string, error) {
	tokenClaims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(j.Expired)).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	return t.SignedString(j.SignKey)
}

func (j *JWTConf) Parse(tokenString string) (interface{}, *JWTClaims, error) {
	tokenClaims := new(JWTClaims)
	token, err := jwt.ParseWithClaims(tokenString, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})
	if err != nil || !token.Valid {
		return nil, nil, ErrInvalidJWT
	}

	return token, tokenClaims, nil
}
