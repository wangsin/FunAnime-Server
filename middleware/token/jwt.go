package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
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
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWT struct {
	SignKey []byte
}

func UserAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			// TODO：完善错误变量处理和日志处理
			ctx.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg": "Token校验失败，非法用户",
			})
			ctx.Abort()
			return
		}

		fmt.Println(token)
		j := NewJWT()
		userInfo, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				ctx.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg": "Token已过期，请重新登录",
				})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg": err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("userInfo", userInfo)
	}
}

func NewJWT() *JWT {
	SecretKey = viper.GetString("secret_key.key")
	if SecretKey == "" {
		// TODO:此处添加日志哈
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