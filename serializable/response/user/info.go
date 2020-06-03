package respUser

import (
	"sinblog.cn/FunAnime-Server/model"
)

type InfoResponse struct {
	Token      string `json:"token,omitempty"`
	UserName   string `json:"user_name"`
	Phone      string `json:"phone"`
	NickName   string `json:"nick_name"`
	Exp        int64  `json:"exp"`
	Level      int    `json:"level"`
	UserId     int64  `json:"user_id"`
	Status     int    `json:"status"`
	Mail       string `json:"mail"`
	UserAvatar string `json:"user_avatar"`
}

func BuildResponse(user *model.User, token string) *InfoResponse {
	return &InfoResponse{
		Token:      token,
		UserAvatar: user.Avatar,
		UserName:   user.Username,
		Phone:      user.Phone,
		NickName:   user.Nickname,
		UserId:     user.Id,
		Status:     user.Status,
		Mail:       user.Mail,
	}
}
