package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	JWTSecretKey  string `mapstructure:"JWT_SRCRET_KEY"`
	DBUrl         string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {

	// 绝对路径，拼接
	path := filepath.Join(GetExecDirectory(), "go-grpc-auth-svc/pkg/config/envs/dev.env")
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

// 获取当前执行程序目录
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err != nil {
		return ""
	}
	return file
}
