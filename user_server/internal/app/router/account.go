package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golearnku/go-practice/user_server/internal/app/api"
	"github.com/golearnku/go-practice/user_server/internal/app/middleware"
)

func init() {
	RouteMgr.register(&AccountRoute{
		Engine: RouteMgr.Engine,
		ctrl:   &api.Account{},
	})
}

type AccountRoute struct {
	*gin.Engine
	ctrl *api.Account
}

func (route AccountRoute) Name() string {
	return "account"
}

// 注册统计路由
func (route AccountRoute) Register() {
	api := route.Group("/v1/account")
	api.Use(middleware.UserAuthMiddleware())
	{
		api.POST("/sms/captcha", route.ctrl.SendSmsCaptcha)
		api.POST("/change_mobile_token", route.ctrl.ChangeMobileToken)
		api.POST("/change_mobile", route.ctrl.ChangeBindingMobile)
	}
}
