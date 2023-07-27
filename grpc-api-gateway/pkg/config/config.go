package config

import "github.com/spf13/viper"

type ports struct {
	GatewayPort     string `mapstructure:"gateway-port"`
	AuthServicePort string `mapstructure:"auth-service-port"`
}

type Config struct {
	PORTS ports
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
