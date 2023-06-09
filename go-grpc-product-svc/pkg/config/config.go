package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	DBUrl         string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {

	// 绝对路径，拼接
	path := filepath.Join("D:/workspace/src/mini-shop-demo/go-grpc-product-svc/pkg", "config/envs/dev.env")
	println(path)
	
	viper.SetConfigFile(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}
