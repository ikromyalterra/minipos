package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func UserGeneratePassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashedPassword)
}

func UserVerifyPassword(userPwd, userStoredPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userStoredPwd), []byte(userPwd))
	return err == nil
}
