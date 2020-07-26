package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golearnku/go-practice/user_server/internal/app/ginplus"
	"github.com/golearnku/go-practice/user_server/internal/app/schema"
	"github.com/golearnku/go-practice/user_server/internal/app/service"
	"github.com/golearnku/go-practice/user_server/pkg/util"
)

type Sms struct {
}

// SendCaptcha 发送验证码
func (sms Sms) SendCaptcha(c *gin.Context) {
	var params schema.SendCaptchaReq
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	// todo: 根据手机号做接口请求限速，防止网络因素重复发送请求和短时间重复发送验证码

	code := util.CreateCaptcha()
	if err := service.NewSms().SendSmsCode(params.Mobile, code); err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResOK(c)
}
