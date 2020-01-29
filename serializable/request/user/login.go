package user

import (
	"github.com/gin-gonic/gin"
)

type LoginRequestInfo struct {
	Phone    string `form:"phone" json:"phone"`
	Password string `form:"password" json:"password"`
	SmsCode  string `form:"smsCode" json:"smsCode"`
}

func (r *LoginRequestInfo) BindRequest(c *gin.Context) error {
	return c.Bind(r)
}

func (r *LoginRequestInfo) CheckRequest() bool {
	if r.Phone == "" || (r.Password == "" && r.SmsCode == "") {
		return false
	} else {
		return true
	}
}
