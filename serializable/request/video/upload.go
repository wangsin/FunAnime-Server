package video

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
)

type UploadRequest struct {
	RemoteId         string `json:"remote_id"`
	CategoryTop      int64  `json:"category_top"`
	CategoryTopDesc  string `json:"category_top_desc"`
	CategoryNext     int64  `json:"category_next"`
	CategoryNextDesc string `json:"category_next_desc"`
	Name             string `json:"name"`
	Desc             string `json:"desc"`
	CoverImg         string `json:"cover_img"`

	UserInfo *token.UserInfo `json:"-"`
}

func (ur *UploadRequest) GetRequest(ctx *gin.Context) error {
	ur.UserInfo = user.GetUserInfoFromContext(ctx)
	err := ctx.Bind(ur)
	if err != nil {
		return err
	}

	return nil
}
