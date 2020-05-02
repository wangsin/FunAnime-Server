package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mervick/aes-everywhere/go/aes256"
	"github.com/spf13/viper"
	"strings"
)

type RegisterRequestInfo struct {
	Phone        string `json:"phone" binding:"required"`
	SmsCode      string `json:"smsCode" binding:"required"`
	Password     string `json:"password"`
	Mail         string `json:"mail"`
	TruePassword string `json:"-"`
}

func (register *RegisterRequestInfo) BindRequest(c *gin.Context) error {
	err := c.Bind(register)
	if err != nil {
		return err
	}

	if register.Mail != "" || register.Password != "" {
		if !strings.Contains(register.Mail, "@") {
			return errors.New("mail_not_fit")
		}

		register.TruePassword = aes256.Decrypt(register.Password, viper.GetString("secret_key.password_key"))
	}

	return nil
}
