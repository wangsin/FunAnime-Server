package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/util/common"
	"sinblog.cn/FunAnime-Server/util/errno"
)

type BasicUser struct {
	UserInfo *token.UserInfo `json:"user_info"`
}

func (bu *BasicUser) GetUserInfo(ctx *gin.Context) error {
	userInfo := GetUserInfoFromContext(ctx)
	if userInfo == nil {
		common.EchoFailedJson(ctx, errno.Uncertified)
		return errors.New("user_not_login")
	}

	bu.UserInfo = userInfo

	return nil
}

func GetUserInfoFromContext(ctx *gin.Context) *token.UserInfo {
	uInfo, ok := ctx.Get("userInfo")
	if !ok {
		//common.EchoJson(ctx, http.StatusOK, errno.Uncertified, nil)
		return nil
	}

	userInfo, ok := uInfo.(*token.UserInfo)
	if !ok {
		//common.EchoFailedJson(ctx, errno.UnknownError)
		return nil
	}

	return userInfo
}