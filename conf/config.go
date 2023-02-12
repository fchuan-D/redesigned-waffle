package conf

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	MysqlUrI  string `mapstructure:"MYSQL_URI"`
	RedisUrI  string `mapstructure:"REDIS_URI"`
	RedisPass string `mapstructure:"REDIS_PASSWORD"`
	RedisDb   int    `mapstructure:"REDIS_DB"`

	Port string `mapstructure:"PORT"`

	JwtKey        string `mapstructure:"JWT_KEY"`
	JwtAccessAge  int    `mapstructure:"JWT_ACCESS_MAXAGE"`
	JwtRefreshAge int    `mapstructure:"JWT_FRESH_MAXAGE"`

	Origin  string `mapstructure:"CLIENT_ORIGIN"`
	BaseUrl string `mapstructure:"BASE_URL"`
}

var config *Config

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	return nil
}

func GetConfig() *Config {
	return config
}

func InitConfig(path string) {
	err := LoadConfig(path)
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
}
