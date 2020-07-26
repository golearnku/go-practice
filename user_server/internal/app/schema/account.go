package schema

// ChangeMobileTokenReq 换绑手机号
type ChangeMobileTokenReq struct {
	Mobile  string `json:"mobile"`
	Captcha string `json:"captcha"`
}

// ChangeBindingMobileReq 换绑手机号
type ChangeBindingMobileReq struct {
	Mobile  string `json:"mobile"`
	Captcha string `json:"captcha"`
	Token   string `json:"token"`
}
