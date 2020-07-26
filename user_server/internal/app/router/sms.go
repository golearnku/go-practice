package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golearnku/go-practice/user_server/internal/app/api"
)

func init() {
	RouteMgr.register(&SmsRoute{
		Engine: RouteMgr.Engine,
		ctrl:   &api.Sms{},
	})
}

type SmsRoute struct {
	*gin.Engine
	ctrl *api.Sms
}

func (route SmsRoute) Name() string {
	return "sms"
}

// 注册路由
func (route SmsRoute) Register() {
	api := route.Group("/v1/sms")
	{
		api.POST("/captcha", route.ctrl.SendCaptcha)
	}
}
