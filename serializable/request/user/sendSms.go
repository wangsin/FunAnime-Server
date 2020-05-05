package user

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	Register = iota + 1 // 注册
	Login               // 登录
)

type SendSmsRequest struct {
	Phone string      `json:"phone"`
	Type  int         `json:"type"`
}

func (register *SendSmsRequest) BindRequest(c *gin.Context) error {
	err := c.Bind(register)
	if err != nil {
		return err
	}

	if register.Phone == "" {
		return errors.New("Params_error")
	}

	return nil
}
