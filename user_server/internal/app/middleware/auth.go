package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golearnku/go-practice/user_server/internal/app/context"
	"github.com/golearnku/go-practice/user_server/internal/app/ginplus"
	"github.com/golearnku/go-practice/user_server/internal/app/injector"
	"github.com/golearnku/go-practice/user_server/pkg/errors"
)

func wrapUserAuthContext(c *gin.Context, userID string) {
	ginplus.SetUserID(c, userID)
	ctx := context.NewUserID(c.Request.Context(), userID)
	c.Request = c.Request.WithContext(ctx)
}

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		userID, err := injector.GetAuther().ParseUserID(c.Request.Context(), ginplus.GetToken(c))
		if err != nil {
			ginplus.ResError(c, errors.ErrInvalidToken)
			return
		}

		wrapUserAuthContext(c, userID)
		c.Next()
	}
}
