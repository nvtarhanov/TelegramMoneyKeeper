package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// It is better to store some of those variables in .env file,
//but for debugging  it is ok
type Config struct {
	Ngrok_url      string `mapstructure:"Ngrok_url"`
	Telegram_token string `mapstructure:"Telegram_token"`
	Port           string `mapstructure:"Port"`
	Telegram_url   string `mapstructure:"Telegram_url"`
	DbConfig       DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `mapstructure:"Host"`
	Username string `mapstructure:"Username"`
	Password string `mapstructure:"Password"`
	Dbname   string `mapstructure:"Dbname"`
	Port     string `mapstructure:"Port"`
	Sslmode  string `mapstructure:"Sslmode"`
}

func NewConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("database", &cfg.DbConfig); err != nil {
		return nil, err
	}

	//Just for log
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed:", e.Name)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Read config error")
		}

	})
	viper.WatchConfig()

	return &cfg, nil
}

func GetConfig() (*Config, error) {
	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("database", &cfg.DbConfig); err != nil {
		return nil, err
	}

	return &cfg, nil
}
