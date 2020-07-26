package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golearnku/go-practice/user_server/internal/app/ginplus"
	"github.com/golearnku/go-practice/user_server/internal/app/injector"
	"github.com/golearnku/go-practice/user_server/internal/app/model"
	"github.com/golearnku/go-practice/user_server/internal/app/schema"
	"github.com/golearnku/go-practice/user_server/internal/app/service"
	"github.com/golearnku/go-practice/user_server/pkg/errors"
	"github.com/golearnku/go-practice/user_server/pkg/util"
)

type Login struct{}

// Login 用户登录
// 没有注册就注册用户)
// 这里需要前端传递端的设备 ID
// todo: 默认用户名生成使用系统的生成，返回一个字段告诉前端是否是新用户，如果是新用户则前端给用户跳转到完善用户信息页面
func (ctrl *Login) Login(c *gin.Context) {
	var (
		err       error
		isNewUser = true
		user      = &model.User{}
		params    schema.LoginReq
	)

	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}
	// todo: Redis 根据手机号加锁，防止网络因素重新请求。

	// todo: 检查验证码是否有效
	if !service.NewSms().CheckSmsCode(params.Mobile, params.Captcha) {
		ginplus.ResError(c, errors.New400Response("验证码错误"))
		return
	}

	// todo: `device_id` 可以防止多个设备批量注册，这里去数据库查询，判断是否超过最大限额

	// todo: 根据手机号判断是否已经注册，已经注册直接登录， 没有注册就走注册逻辑
	if isNewUser {
		userName := util.GetUsername()
		user, err = service.NewAccount().Register(userName, params.Mobile, params.DeviceID)
		if err != nil {
			ginplus.ResError(c, errors.New400Response("注册用户失败"))
			return
		}
	}

	userID := user.ID
	// 将用户ID放入上下文
	ginplus.SetUserID(c, userID)
	srv := service.NewLogin(injector.GetAuther())
	tokenInfo, err := srv.GenerateToken(c.Request.Context(), userID)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResSuccess(c, tokenInfo)
}

// Logout 用户登出
func (ctrl *Login) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	// 检查用户是否处于登录状态，如果是则执行销毁
	userID := ginplus.GetUserID(c)
	if userID != "" {
		err := service.NewLogin(injector.GetAuther()).DestroyToken(ctx, ginplus.GetToken(c))
		if err != nil {
			ginplus.ResError(c, err)
			return
		}
	}
	ginplus.ResOK(c)
}

// RefreshToken 刷新令牌
func (ctrl *Login) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	tokenInfo, err := service.NewLogin(injector.GetAuther()).GenerateToken(ctx, ginplus.GetUserID(c))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, tokenInfo)
}
