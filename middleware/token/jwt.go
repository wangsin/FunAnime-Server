package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"sinblog.cn/FunAnime-Server/util/common"
	"sinblog.cn/FunAnime-Server/util/errno"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
	SecretKey        string
)

type UserInfo struct {
	UserId   int64  `json:"userId"`
	Level    int    `json:"level"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
	Sex      int8   `json:"male"`
	jwt.StandardClaims
}

func (user *UserInfo) ToJson() string {
	str, _ := json.Marshal(user)
	return string(str)
}

type JWT struct {
	SignKey []byte
}

func ParseToken(ctx *gin.Context) *UserInfo {
	token := ctx.Request.Header.Get("token")
	if token == "" {
		common.EchoFailedJson(ctx, errno.TokenInvalid)
		return nil
	}

	j := NewJWT()
	userInfo, err := j.ParseToken(token)
	if err != nil {
		if err == TokenExpired {
			common.EchoFailedJson(ctx, errno.TokenExpired)
			return nil
		}
		common.EchoFailedJson(ctx, errno.UnknownError)
		return nil
	}

	return userInfo
}

func NewJWT() *JWT {
	SecretKey = viper.GetString("secret_key.key")
	if SecretKey == "" {
		panic(fmt.Errorf("Read Config File Failed At Init Token\n"))
	}

	return &JWT{
		SignKey: []byte(SecretKey),
	}
}

func (j *JWT) ParseToken(tokenString string) (*UserInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserInfo{}, func(t *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			fmt.Println(ve.Errors)
			switch ve.Errors {
			case jwt.ValidationErrorMalformed:
				return nil, TokenMalformed
			case jwt.ValidationErrorExpired:
				return nil, TokenExpired
			case jwt.ValidationErrorNotValidYet:
				return nil, TokenNotValidYet
			default:
				return nil, TokenInvalid
			}
		}
	}

	if userInfo, ok := token.Claims.(*UserInfo); ok && token.Valid {
		return userInfo, nil
	}

	return nil, TokenInvalid
}

func (j *JWT) CreateToken(userInfo *UserInfo) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userInfo)
	return token.SignedString(j.SignKey)
}
