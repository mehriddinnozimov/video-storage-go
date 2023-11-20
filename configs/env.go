package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Envs struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DATABASE_HOST"`
	DBPort         string `mapstructure:"DATABASE_PORT"`
	DBUser         string `mapstructure:"DATABASE_USER"`
	DBPassword     string `mapstructure:"DATABASE_PASSWORD"`
	DB             string `mapstructure:"DATABASE"`
	JwtSecret      string `mapstructure:"JWT_SECRET"`
	EskizLogin     string `mapstruckture:"ESKIZ_LOGIN"`
	EskizPassword  string `mapstruckture:"ESKIZ_PASSWORD"`
	PspSecretKey   string `mapstruckture:"PSP_SECRET_KEY"`
	PspVendorId    int    `mapstruckture:"PSP_VENDOR_ID"`
	PaymentUrl     string `mapstruckture:"PAYMENT_URL"`
	GinMode        string `mapstruckture:"GIN_MODE"`
	Domain         string `mapstruckture:"DOMAIN"`
	IsSecure       bool   `mapstruckture:"IS_SECURE"`
}

func GetEnv() *Envs {
	env := Envs{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	viper.Unmarshal(&env)

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
