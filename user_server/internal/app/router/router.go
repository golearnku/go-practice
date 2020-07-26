package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golearnku/go-practice/user_server/internal/app/middleware"
)

var RouteMgr = NewRouteManager()

type Router interface {
	Name() string
	Register()
}

type RouteManager struct {
	Engine  *gin.Engine
	handler map[string]Router
}

func NewRouteManager() *RouteManager {
	routeMgr := &RouteManager{
		Engine:  gin.New(),
		handler: make(map[string]Router),
	}
	routeMgr.Engine.NoMethod(middleware.NoMethodHandler())
	routeMgr.Engine.NoRoute(middleware.NoRouteHandler())
	return routeMgr
}

func (route *RouteManager) Load() {
	route.middleware()
	for _, r := range route.handler {
		r.Register()
	}
}

func (route *RouteManager) middleware() {
	// 崩溃恢复
	route.Engine.Use(middleware.RecoveryMiddleware())
}

func (route *RouteManager) register(router Router) {
	route.handler[router.Name()] = router
}
