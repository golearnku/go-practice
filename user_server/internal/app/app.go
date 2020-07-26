package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golearnku/go-practice/user_server/internal/app/config"
	"github.com/golearnku/go-practice/user_server/internal/app/injector"
	. "github.com/golearnku/go-practice/user_server/internal/app/router"
	"github.com/golearnku/go-practice/user_server/pkg/util"
	logger "github.com/golearnku/sdk-zap"
	"go.uber.org/zap"
)

type options struct {
	Env        string
	ConfigFile string
	WWWDir     string
	Version    string
}

// Option 定义配置项
type Option func(*options)

// SetEnv 设置环境变量
func SetEnv(env string) Option {
	return func(o *options) {
		o.Env = env
	}
}

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetWWWDir 设定静态站点目录
func SetWWWDir(s string) Option {
	return func(o *options) {
		o.WWWDir = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

// InitGinEngine 初始化 gin
func InitGinEngine() *gin.Engine {
	gin.SetMode(config.C.RunMode)
	RouteMgr.Load()
	app := RouteMgr.Engine
	return app
}

// InitHTTPServer 初始化http服务
func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	cfg := config.C.HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.Logger().Info("HTTP server is running", zap.String("addr", addr))
		var err error
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Logger().Error("Shutdown", zap.Error(err))
		}
	}
}

// Init 应用初始化
func Init(ctx context.Context, opts ...Option) (func(), error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	if env := util.GetEnv(); env != "" {
		o.Env = env
	}
	if confPath := util.GetConfPath(); confPath != "" {
		o.ConfigFile = confPath
	}
	if err := config.MustLoad(config.SetEnv(o.Env), config.SetPath(o.ConfigFile), config.SetPrint(true)); err != nil {
		return nil, err
	}
	if v := o.WWWDir; v != "" {
		config.C.WWW = v
	}

	// 初始化日志模块
	logger.New(logger.SetPath(config.C.Log.Path), logger.SetEnv(o.Env), logger.SetDebug(config.C.Log.Debug), logger.SetOutput(config.C.Log.Output))

	logger.Logger().Info("服务已启动", zap.String("运行模式", o.Env), zap.String("版本号", o.Version), zap.Int("进程号", os.Getpid()))

	// 初始化 gin 服务
	engine := InitGinEngine()

	// 初始化HTTP服务
	httpServerCleanFunc := InitHTTPServer(ctx, engine)

	// 初始化 jwt
	_, err := injector.InitAuth()
	if err != nil {
		return nil, err
	}

	return func() {
		httpServerCleanFunc()
	}, nil
}

// Run 运行服务
func Run(ctx context.Context, opts ...Option) error {
	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx, opts...)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logger.Logger().Info("接收到关闭指令", zap.String("信号", sig.String()))
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.CompareAndSwapInt32(&state, 1, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logger.Logger().Info("服务退出")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
	return nil
}
