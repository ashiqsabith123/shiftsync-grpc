package config

import "github.com/spf13/viper"

type dbConfig struct {
	DbHost     string `mapstructure:"host"`
	DbName     string `mapstructure:"dbname"`
	DbPort     string `mapstructure:"port"`
	DbUser     string `mapstructure:"user"`
	DbPaswword string `mapstructure:"password"`
}

type CryptoConfig struct {
	Secret string `mapstructure:"secret"`
}

type port struct {
	SvcPort string `mapstructure:"port"`
}

type Config struct {
	Db     dbConfig     `mapstructure:"db"`
	Port   port         `mapstructure:"svc-port"`
	Crypto CryptoConfig `mapstructure:"crypto"`
}

var config Config

func LoadConfig() (Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("pkg/config/")

	if err := vp.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := vp.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetCryptoSecret() string {
	return config.Crypto.Secret
}
