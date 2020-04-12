package user

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/cache"
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/util/common"
	"sinblog.cn/FunAnime-Server/util/errno"
)

func UserAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.Request.Header.Get("token")
		if tokenStr == "" {
			common.EchoFailedJson(ctx, errno.TokenInvalid)
			return
		}

		j := token.NewJWT()
		userInfo, err := j.ParseToken(tokenStr)
		if err != nil {
			if err == token.TokenExpired {
				common.EchoFailedJson(ctx, errno.TokenExpired)
				return
			}
			common.EchoFailedJson(ctx, errno.UnknownError)
			return
		}

		_, err = cache.GetUserLogin(userInfo.UserId)
		if err != nil {
			common.EchoFailedJson(ctx, errno.TokenExpired)
			return
		}

		ctx.Set("userInfo", userInfo)
	}
}

