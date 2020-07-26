package util

import (
	"os"
)

var (
	_env      string
	_confPath string
)

func init() {
	_env = os.Getenv("FRAME_ENV")
	_confPath = os.Getenv("FRAME_CONF_PATH")
}

// GetEnv 获取程序运行环境变量
func GetEnv() string {
	return _env
}

// GetConfPath 获取配置文件路径
func GetConfPath() string {
	return _confPath
}
