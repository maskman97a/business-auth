package hsm

import "github.com/spf13/viper"

type Config struct {
	Url           string
	SignEndpoint  string
	VerifyEnpoint string
}

func NewHsmConfig() *Config {
	return &Config{Url: viper.GetString("HSM_URL"),
		SignEndpoint:  viper.GetString("HSM_SIGN_ENDPOINT"),
		VerifyEnpoint: viper.GetString("HSM_VERIFY_ENDPOINT")}
}
