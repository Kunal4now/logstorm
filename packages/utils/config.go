package utils

import "github.com/spf13/viper"

type Config struct {
	HOST     string `mapstructure:"HOST"`
	PORT     int `mapstructure:"PORT"`
	PASSWORD string `mapstructure:"PASSWORD"`
	DBNAME   string `mapstructure:"DBNAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return 
}
