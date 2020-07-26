package model

import (
	"time"
)

// User 用户对象
type User struct {
	ID        string    // ID 自增ID
	DeviceID  string    // DeviceID 设备号
	Username  string    // Username 用户名
	Mobile    string    // Mobile 手机号
	Status    int       // Status 用户状态
	CreatedAt time.Time // CreatedAt 创建时间
	UpdatedAt time.Time // UpdatedAt 更新时间
}

func NewUser() *User {
	return &User{}
}
