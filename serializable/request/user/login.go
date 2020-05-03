package user

import (
	"github.com/gin-gonic/gin"
)

type LoginRequestInfo struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	SmsCode  string `json:"smsCode"`
}

func (r *LoginRequestInfo) BindRequest(c *gin.Context) error {
	return c.Bind(r)
}

func (r *LoginRequestInfo) CheckRequest() bool {
	if (r.Phone == "" && r.Password == "") || (r.Phone == "" && r.SmsCode == "") {
		return false
	} else {
		return true
	}
}
