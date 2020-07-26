package service

import (
	"github.com/golearnku/go-practice/user_server/pkg/errors"
)

var (
	SendSmsExceedingLimit = errors.New400Response("发送短信超出限制")
)

type Sms struct{}

func NewSms() *Sms {
	return &Sms{}
}

// SendSmsCode 发送验证码
func (srv *Sms) SendSmsCode(mobile, code string) (err error) {
	// fixme 这里基于 Redis 的 hash 简单增加了恶意发送短信限制和 每天短信限制，还可以基于 ip 等防护
	if err = srv.checkSendMobileDailyLimit(mobile); err != nil {
		return SendSmsExceedingLimit
	}
	if err = srv.checkSendMobileMinuteLimit(mobile); err != nil {
		return SendSmsExceedingLimit
	}

	// fixme: 调用短信网关发送短信

	// fixme: 记录网关返回内容，和错误信息等。

	// fixme: redis string set 缓存验证码有效期 10 分钟

	// fixme: redis string set 记录 2 分钟频率 ，2 分钟过期

	// fixme: redis string set 记录 1 天频率
	// expiration := util.LastSecondOfDay(time.Now()).Unix() - time.Now().Unix() // 这是一天的redis 过期时间

	return nil
}

// CheckSmsCode 检查验证码是否有效
func (srv Sms) CheckSmsCode(mobile, code string) bool {
	// fixme: 验证码验证逻辑，判断验证码是否有效
	return true
}

// checkSendMobileMinuteLimit 检查手机号是否 超过 2 分钟限额
func (srv *Sms) checkSendMobileMinuteLimit(mobile string) error {
	// fixme 基于 Redis 的 get 获取 这个手机号
	return nil
}

// checkSendMobileDailyLimit 检查手机号是否超过每天的限额
func (srv *Sms) checkSendMobileDailyLimit(mobile string) error {
	// fixme 基于 Redis 的 get 获取 这个手机号
	return nil
}
