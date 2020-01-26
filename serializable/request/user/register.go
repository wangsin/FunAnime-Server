package user

import "github.com/gin-gonic/gin"

type RegisterRequestInfo struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
}

func (register *RegisterRequestInfo) BindRequest(c *gin.Context) error {
	return c.Bind(register)
}
