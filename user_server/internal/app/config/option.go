package config

type options struct {
	env   string
	path  string
	print bool
}

type Jwt struct {
	Enable        bool   `mapstructure:"enable"`
	SigningMethod string `mapstructure:"signing_method"`
	SigningKey    string `mapstructure:"signing_key"`
	Expired       int    `mapstructure:"expired"`
	Store         string `mapstructure:"store"`
	FilePath      string `mapstructure:"file_path"`
	RedisDB       int    `mapstructure:"redis_db"`
	RedisPrefix   string `mapstructure:"redis_prefix"`
}

// HTTP http配置参数
type HTTP struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	CertFile        string `mapstructure:"cert_file"`
	KeyFile         string `mapstructure:"key_file"`
	ShutdownTimeout int    `mapstructure:"shutdown_timeout"`
}

// Log 日志配置参数
type Log struct {
	Debug  bool   `mapstructure:"debug"`
	Output bool   `mapstructure:"output"`
	Path   string `mapstructure:"path"`
}

// Option 配置选项
type Option func(*options)

// SetEnv 设置环境变量
func SetEnv(env string) Option {
	return func(o *options) {
		o.env = env
	}
}

// SetPath 设置配置环境路径
func SetPath(path string) Option {
	return func(o *options) {
		o.path = path
	}
}

// SetPrint 设置是否打印配置信息
func SetPrint(print bool) Option {
	return func(o *options) {
		o.print = print
	}
}
