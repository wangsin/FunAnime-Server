package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RequestInfo struct {
	Phone    string `form:"phone" json:"phone"`
	Mail     string `form:"mail" json:"mail"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (r *RequestInfo) BindRequest(c *gin.Context) error {
	return c.Bind(r)
}

func (r *RequestInfo) CheckRequest() bool {
	fmt.Println(*r)
	if r.Password == "" {
		return false
	} else {
		if r.Mail == "" && r.Phone == "" && r.Username == "" {
			return false
		}
		return true
	}
}
