package user

import "github.com/gin-gonic/gin"

const (
	Register = iota + 1
	Login
)

type SendSmsRequest struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
	Type  int    `form:"type" json:"type" binding:"required"`
}

func (register *SendSmsRequest) BindRequest(c *gin.Context) error {
	return c.Bind(register)
}
