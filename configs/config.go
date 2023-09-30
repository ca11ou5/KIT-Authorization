package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DBHOST"`
	DBUser     string `mapstructure:"DBUSER"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	DBName     string `mapstructure:"DBNAME"`
	DBPort     string `mapstructure:"DBPORT"`
	DBSSLMode  string `mapstructure:"DBSSLMODE"`
}

func InitConfig() (c Config) {
	viper.SetConfigFile("configs/envs/dev.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read config")
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatal("failed to unmarshal config")
	}

	return
}
