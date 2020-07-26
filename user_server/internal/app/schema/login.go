package schema

// LoginReq 登录
type LoginReq struct {
	DeviceID string `json:"device_id"` // 设备 ID
	Mobile   string `json:"mobile"`    // 手机号
	Captcha  string `json:"captcha"`   // 验证码
}

// LoginTokenInfo 登录令牌信息
type LoginTokenInfo struct {
	AccessToken string `json:"access_token"` // 访问令牌
	TokenType   string `json:"token_type"`   // 令牌类型
	ExpiresAt   int64  `json:"expires_at"`   // 令牌到期时间戳
}
