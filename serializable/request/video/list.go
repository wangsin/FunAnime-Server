package video

import (
	"github.com/gin-gonic/gin"
)

type GetVideoListForOuter struct {
	Category int `form:"category"`
	Page     int `form:"page"`
	Size     int `form:"size"`
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