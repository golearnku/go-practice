package schema

import (
	"time"
)

const (
	SmsLockKey            = "sms|mobile|"
	SmsCaptcha            = "sms|captcha|"
	SmsIntervalLock       = "sms|sendlimit|"
	SmsDayLock            = "sms|daylimit|"
	SmsCaptchaExpire      = time.Minute * 10
	SmsIntervalLockExpire = time.Minute * 2
)

// SendCaptchaReq 发送验证码
type SendCaptchaReq struct {
	Mobile string `json:"mobile"`
}
