package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)
type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
}

func LoadConfig() (c Config, err error) {

	// 绝对路径，拼接
	path := filepath.Join("D:/workspace/src/mini-shop-demo/", "go-grpc-api-gateway/pkg/config/envs/dev.env")

	viper.SetConfigFile(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}

// 获取当前执行程序目录
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err != nil {
		return ""
	}
	return file
}
