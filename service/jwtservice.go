package service

import (
	"time"

	"cpipi1024.com/minicloud/bootstrap"
	"cpipi1024.com/minicloud/db"
	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtkey = bootstrap.SrvConf.JWT.SecretKey //签名私钥
	ttl    = bootstrap.SrvConf.JWT.TTL       // ttl
)

// token数据
type TokenData struct {
	TokenType string `json:"token_type"` // token类型
	TokenStr  string `json:"token_str"`  // token字符串
	ExpiredIn int    `json:"expired_in"` // token过期时间
}

// 自定义载荷
type CustomClaim struct {
	UserName string      `json:"user_name"` // 用户名
	UUID     string      `json:"uuid"`      // 用户uuid
	Role     db.RoleType `json:"role"`      // 用户角色
	jwt.RegisteredClaims
}

// jwtuser接口
type JwtUser interface {
	GetUUID() string
	GetRole() db.RoleType
	GetName() string
}

type jwtService struct{}

var JwtService = new(jwtService)

// todo: 创建token
func (service *jwtService) CreateToken(user JwtUser) (*jwt.Token, TokenData, error) {
	custClaim := CustomClaim{
		user.GetName(),
		user.GetUUID(),
		user.GetRole(),
		jwt.RegisteredClaims{
			ID:        user.GetUUID(),                                                       // jwtid
			Issuer:    "minicloud",                                                          // 签发人
			NotBefore: jwt.NewNumericDate(time.Now().Add(-5 * time.Second)),                 // 生效前
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttl) * time.Second)), // 失效
		},
	}

	// 创建token对象 包含负载内容
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, custClaim)

	// 使用私钥进行签名
	str, err := token.SignedString(jwtkey)

	if err != nil {
		return nil, TokenData{}, err
	}

	data := TokenData{
		TokenType: "bearer",
		TokenStr:  str,
		ExpiredIn: ttl,
	}

	return token, data, nil

}
