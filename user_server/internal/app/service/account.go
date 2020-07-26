package service

import (
	"time"

	"github.com/golearnku/go-practice/user_server/internal/app/model"
	"github.com/golearnku/go-practice/user_server/pkg/util"
)

type Account struct {
}

func NewAccount() *Account {
	return &Account{}
}

//  Register 注册用户
func (srv Account) Register(username, mobile, deviceID string) (user *model.User, err error) {
	// fixme: 保存用户信息到数据库
	user = &model.User{
		ID:        util.NewUUID(),
		Username:  username,
		DeviceID:  deviceID,
		Mobile:    mobile,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return user, nil
}
