package user

import "github.com/gin-gonic/gin"

type SendSmsRequest struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
}

func (register *SendSmsRequest) BindRequest(c *gin.Context) error {
	return c.Bind(register)
}

