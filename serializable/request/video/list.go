package video

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
)

type GetVideoListForOuter struct {
	Category int    `form:"category"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	Title    string `form:"title"`
}

func (req *GetVideoListForOuter) GetRequest(c *gin.Context) error {
	err := c.Bind(req)
	if err != nil {
		return err
	}

	if req.Size == 0 {
		req.Size = 12
	}

	if req.Page == 0 {
		req.Page = 1
	}

	return nil
}

type GetVideoListForCollection struct {
	Page     int             `form:"page"`
	Size     int             `form:"size"`
	Name     string          `form:"name"`
	UserInfo *token.UserInfo `json:"-"`
}

func (req *GetVideoListForCollection) GetRequest(c *gin.Context) error {
	req.UserInfo = user.GetUserInfoFromContext(c)
	err := c.Bind(req)
	if err != nil {
		return err
	}

	if req.Size == 0 {
		req.Size = 12
	}

	if req.Page == 0 {
		req.Page = 1
	}

	return nil
}