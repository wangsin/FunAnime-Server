package video

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
)

type CreateCollection struct {
	VideoId  int64           `json:"video_id"`
	UserInfo *token.UserInfo `json:"-"`
}

func (req *CreateCollection) GetRequest(c *gin.Context) error {
	req.UserInfo = user.GetUserInfoFromContext(c)
	err := c.Bind(req)
	if err != nil {
		return err
	}

	return nil
}

type RemoveCollection struct {
	VideoId  int64           `json:"video_id"`
	UserInfo *token.UserInfo `json:"-"`
}

func (req *RemoveCollection) GetRequest(c *gin.Context) error {
	req.UserInfo = user.GetUserInfoFromContext(c)
	err := c.Bind(req)
	if err != nil {
		return err
	}

	return nil
}

type UpdateVideoInfo struct {
	VideoId   int64           `json:"video_id"`
	VideoName string          `json:"video_name"`
	VideoDesc string          `json:"video_desc"`
	UserInfo  *token.UserInfo `json:"-"`
}

func (req *UpdateVideoInfo) GetRequest(c *gin.Context) error {
	req.UserInfo = user.GetUserInfoFromContext(c)
	err := c.Bind(req)
	if err != nil {
		return err
	}

	return nil
}
