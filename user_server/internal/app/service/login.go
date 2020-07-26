package service

import (
	"context"

	"github.com/golearnku/go-practice/user_server/internal/app/schema"
	"github.com/golearnku/go-practice/user_server/pkg/auth"
	"github.com/golearnku/go-practice/user_server/pkg/errors"
)

type Login struct {
	Auth auth.Auther
}

func NewLogin(auth auth.Auther) *Login {
	return &Login{Auth: auth}
}

// GenerateToken 生成令牌
func (srv *Login) GenerateToken(ctx context.Context, userID string) (*schema.LoginTokenInfo, error) {
	tokenInfo, err := srv.Auth.GenerateToken(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	item := &schema.LoginTokenInfo{
		AccessToken: tokenInfo.GetAccessToken(),
		TokenType:   tokenInfo.GetTokenType(),
		ExpiresAt:   tokenInfo.GetExpiresAt(),
	}
	return item, nil
}

// DestroyToken 销毁令牌
func (srv *Login) DestroyToken(ctx context.Context, tokenString string) error {
	err := srv.Auth.DestroyToken(ctx, tokenString)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
