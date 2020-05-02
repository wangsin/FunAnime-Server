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
	Ctx   gin.Context `json:"-"`
}

func (register *SendSmsRequest) BindRequest(c *gin.Context) error {
	register.Ctx = *c
	err := c.Bind(register)
	if err != nil {
		return err
	}

	if register.Phone == "" {
		return errors.New("Params_error")
	}

	return nil
}
