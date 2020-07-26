package main

import (
	"context"
	"os"

	"github.com/golearnku/go-practice/user_server/internal/app"
	"github.com/urfave/cli/v2"
)

// VERSION 版本号，可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "1.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "user-server"
	app.Version = VERSION
	app.Usage = "用户服务"
	app.Commands = []*cli.Command{
		newWebCmd(context.Background()),
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "运行web服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "env",
				Aliases:     []string{"e"},
				Usage:       "环境变量",
				DefaultText: "dev",
			},
			&cli.StringFlag{
				Name:  "www",
				Usage: "静态站点目录",
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetWWWDir(c.String("www")),
				app.SetEnv(c.String("env")),
				app.SetVersion(VERSION))
		},
	}
}
