package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golearnku/go-practice/user_server/internal/app/ginplus"
	"github.com/golearnku/go-practice/user_server/internal/app/schema"
	"github.com/golearnku/go-practice/user_server/internal/app/service"
	"github.com/golearnku/go-practice/user_server/pkg/errors"
	"github.com/golearnku/go-practice/user_server/pkg/util"
)

type Account struct {
}

// SendSmsCaptcha 账户原手机号发送验证码
func (account Account) SendSmsCaptcha(c *gin.Context) {
	var params schema.SendCaptchaReq
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	// todo: 根据手机号做接口请求限速，防止网络因素重复发送请求，和 短时间重复发送验证码

	// todo 验证输入的手机号是否和账号绑定的手机号一致
	var mobile string
	if params.Mobile != mobile {
		ginplus.ResError(c, errors.New400Response("你输入的手机号和账户绑定的手机号不一致"))
		return
	}

	// 短信网关，发送验证码
	code := util.CreateCaptcha()
	if err := service.NewSms().SendSmsCode(params.Mobile, code); err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResOK(c)
}

// ChangeMobileToken 换绑手机号，旧手机号获取token
func (account Account) ChangeMobileToken(c *gin.Context) {
	var params schema.ChangeMobileTokenReq
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	// todo 验证是否绑定手机，没有绑定手机则不能换绑

	// todo 验证验证码是否正确
	if !service.NewSms().CheckSmsCode(params.Mobile, params.Captcha) {
		ginplus.ResError(c, errors.New400Response("验证码错误"))
		return
	}

	token := util.NewUUID()

	// todo  存储到 Redis string set  10 分钟过期

	ginplus.ResSuccess(c, map[string]string{
		"token": token,
	})
}

// ChangeBindingMobile 换绑手机号
func (account Account) ChangeBindingMobile(c *gin.Context) {
	var params schema.ChangeBindingMobileReq
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	// todo 绑定的手机号没有更换

	// todo 验证token是否正确

	// todo 验证新手机号验证码是否正确

	// todo 判断新手机号是否已经被绑定

	// todo 更新数据库手机号
	ginplus.ResOK(c)
}
