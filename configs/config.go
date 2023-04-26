package configs

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port string
}

type MongoDBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type LogConfig struct {
	Level      string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}
type Config struct {
	Server  AppConfig
	MongoDB MongoDBConfig
	Redis   RedisConfig
	Log     LogConfig
}

var (
	instance *Config
	once     sync.Once
)

func init() {
	once.Do(func() {
		viper.SetConfigName("./configs/config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}

		instance = &Config{}
		if err := viper.Unmarshal(instance); err != nil {
			log.Fatalf("Failed to unmarshal config: %v", err)
		}
	})
}

func GetConfig() *Config {
	return instance
}
