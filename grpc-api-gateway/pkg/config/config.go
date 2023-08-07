package config

import "github.com/spf13/viper"

type ports struct {
	GatewayPort      string `mapstructure:"gateway-port"`
	AuthServicePort  string `mapstructure:"auth-service-port"`
	PunchServicePort string `mapstructure:"punch-service-port"`
	FormServicePort string `mapstructure:"form-service-port"`
}

type JWTConfig struct {
	Secret_key string `mapstructure:"secret_key"`
}

type Config struct {
	PORTS ports     `mapstructure:"port"`
	JWT   JWTConfig `mapstructure:"jwt"`
	Db    DBConfig  `mapstructure:"db"`
}

type DBConfig struct {
	DbHost     string `mapstructure:"host"`
	DbName     string `mapstructure:"dbname"`
	DbPort     string `mapstructure:"port"`
	DbUser     string `mapstructure:"user"`
	DbPaswword string `mapstructure:"password"`
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

func JwtConfig() string {
	return config.JWT.Secret_key
}
