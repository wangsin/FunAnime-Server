package respUser

import (
	"sinblog.cn/FunAnime-Server/model"
)

type InfoResponse struct {
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	NickName string `json:"nick_name"`
	Exp      int64  `json:"exp"`
	Level    int    `json:"level"`
	UserId   int64  `json:"user_id"`
}

func BuildResponse(user *model.User) *InfoResponse {
	return &InfoResponse{
		UserName: user.Username,
		Phone:    user.Phone,
		NickName: user.Nickname,
		Exp:      user.ExpCount,
		Level:    user.Level,
		UserId:   user.Id,
	}
}