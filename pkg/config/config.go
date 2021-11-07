package config

import "github.com/spf13/viper"

// It is better to store some of those variables in .env file,
//but for debugging  it is ok
type Config struct {
	Ngrok_url      string `mapstructure:"Ngrok_url"`
	Telegram_token string `mapstructure:"Telegram_token"`
	Port           string `mapstructure:"Port"`
	Telegram_url   string `mapstructure:"Telegram_url"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
