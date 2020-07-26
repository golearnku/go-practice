package injector

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/golearnku/go-practice/user_server/internal/app/config"
	"github.com/golearnku/go-practice/user_server/pkg/auth"
	"github.com/golearnku/go-practice/user_server/pkg/auth/jwtauth"
	"github.com/golearnku/go-practice/user_server/pkg/auth/jwtauth/store/buntdb"
)

var (
	initAuther auth.Auther
)

// InitAuth 初始化用户认证
func InitAuth() (auth.Auther, error) {
	cfg := config.C.Jwt
	var opts []jwtauth.Option
	opts = append(opts, jwtauth.SetExpired(cfg.Expired))
	opts = append(opts, jwtauth.SetSigningKey([]byte(cfg.SigningKey)))
	opts = append(opts, jwtauth.SetKeyfunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(cfg.SigningKey), nil
	}))

	var method jwt.SigningMethod
	switch cfg.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}
	opts = append(opts, jwtauth.SetSigningMethod(method))

	var store jwtauth.Storer
	s, err := buntdb.NewStore(cfg.FilePath)
	if err != nil {
		return nil, err
	}
	store = s

	initAuther = jwtauth.New(store, opts...)
	return initAuther, nil
}

func GetAuther() auth.Auther {
	if initAuther == nil {
		panic("auther is not initialized")
	}
	return initAuther
}
