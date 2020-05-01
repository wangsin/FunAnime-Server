package user

import "github.com/gin-gonic/gin"

const (
	Register = iota + 1 // 注册
	Login               // 登录
)

type SendSmsRequest struct {
	Phone string      `json:"phone" binding:"required"`
	Type  int         `json:"type" binding:"required"`
	Ctx   gin.Context `json:"-"`
}

func (register *SendSmsRequest) BindRequest(c *gin.Context) error {
	register.Ctx = *c
	return c.Bind(register)
}
