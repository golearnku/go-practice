package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/golearnku/go-practice/user_server/pkg/errors"
	"github.com/golearnku/go-practice/user_server/pkg/util"
	"github.com/spf13/viper"
)

var (
	// C 全局配置(需要先执行MustLoad，否则拿不到配置)
	C = new(Config)
)

// Config 配置参数
type Config struct {
	RunMode string `mapstructure:"run_mode"`
	WWW     string `mapstructure:"www"`
	Jwt     Jwt    `mapstructure:"jwt"`
	Log     Log    `mapstructure:"log"`
	HTTP    HTTP   `mapstructure:"http"`
	opts    options
}

// MustLoad 加载解析配置信息
func MustLoad(opts ...Option) (err error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	C.opts = o
	if o.env == "" || o.path == "" {
		return errors.New("请设置路径和环境变量")
	}
	viper.SetConfigFile(filepath.Join(o.path, o.env+".config.toml"))
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("USER_SERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	if err = viper.Unmarshal(C); err != nil {
		return err
	}
	C.PrintJSON()
	return nil
}

// PrintJSON 基于JSON格式输出配置
func (c Config) PrintJSON() {
	if !c.opts.print {
		return
	}
	buf, err := util.JSONMarshalIndent(&C, "", "")
	if err != nil {
		os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
		return
	}
	os.Stdout.WriteString(string(buf) + "\n")
}

// GetEnv 获取环境变量
func (c Config) GetEnv() string {
	return c.opts.env
}

// GetCofPath 获取配置文件跟路径
func (c Config) GetCofPath() string {
	return c.opts.path
}
