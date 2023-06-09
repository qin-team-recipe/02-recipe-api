package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ApplicationName string `mapstructure:"APPLICATION_NAME"`

	ServerPort          string `mapstructure:"SERVER_PORT"`
	ContainerServerPort string `mapstructure:"CONTAINER_SERVER_PORT"`
	Environment         string `mapstructure:"ENV"`
	// database
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`

	TokenExpireAt time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	SecretKey     string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`

	GoogleClientID  string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleSecretKey string `mapstructure:"GOOGLE_SECRET_KEY"`

	Cors struct {
		AllowOrigins []string `mapstructure:"CORS_ALLOW_ORINGINS"`
	}

	// Jwt struct {
	// TokenExpireAt int    `mapstructure:"JWT_TOKEN_EXPIRE_AT"`
	// SecretKey     string `mapstructure:"JWT_SECRET_KEY"`
	// }
}

func NewConfig(path string) *Config {
	viper.SetConfigFile(path)
	viper.SetConfigFile(".env")
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error read config: ", err)
	}

	c := &Config{}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal("env file unmarshal: ", err)
	}

	if c.Environment == "development" {
		log.Print("\n==================================================\n")
		log.Print("\n\nCurrently, it's a development environment!!!!!!\n\n")
		log.Print("\n==================================================\n")
	}

	c.Cors.AllowOrigins = viper.GetStringSlice("CORS_ALLOW_ORIGINS")

	return c
}
