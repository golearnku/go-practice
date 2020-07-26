package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golearnku/go-practice/user_server/internal/app/api"
	"github.com/golearnku/go-practice/user_server/internal/app/middleware"
)

func init() {
	RouteMgr.register(&LoginRoute{
		Engine: RouteMgr.Engine,
		ctrl:   &api.Login{},
	})
}

type LoginRoute struct {
	*gin.Engine
	ctrl *api.Login
}

func (route LoginRoute) Name() string {
	return "login"
}

// 注册路由
func (route LoginRoute) Register() {
	api := route.Group("/v1")
	api.Use(middleware.UserAuthMiddleware(middleware.AllowPathPrefixSkipper("/v1/login")))
	{
		api.POST("/login", route.ctrl.Login)
		api.DELETE("/logout", route.ctrl.Logout)
		api.POST("/current/refresh-token", route.ctrl.RefreshToken)
	}
}
