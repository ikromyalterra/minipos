package config

import (
	"github.com/spf13/viper"
)

type (
	// JWTConfig JWTConfig
	JWTConfig struct {
		SignKey []byte
		Expired int
	}
)

func GetJWTSigKey() []byte {
	return []byte(viper.GetString("jwt.sign_key"))
}

func GetJWTExpired() int {
	return viper.GetInt("jwt.expired_in_hour")
}
