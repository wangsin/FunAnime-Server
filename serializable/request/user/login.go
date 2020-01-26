package user

import (
	"github.com/gin-gonic/gin"
)

type LoginRequestInfo struct {
	Phone    string `form:"phone" json:"phone"`
	Mail     string `form:"mail" json:"mail"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (r *LoginRequestInfo) BindRequest(c *gin.Context) error {
	return c.Bind(r)
}

func (r *LoginRequestInfo) CheckRequest() bool {
	if r.Password == "" {
		return false
	} else {
		if r.Mail == "" && r.Phone == "" && r.Username == "" {
			return false
		}
		return true
	}
}
