package user

import (
	"github.com/jinzhu/gorm"
	"sinblog.cn/FunAnime-Server/model"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
	"sinblog.cn/FunAnime-Server/util/errno"
)

func RegisterUser(userRequest user.RegisterRequestInfo) int64 {
	_, userCount, err := model.QueryUserWithWhereMap(
		map[string]interface{}{
			"phone": userRequest.Phone,
		},
		map[string]interface{}{
			"status <> ?": model.UserDeleted,
		},
	)

	if err != nil && err != gorm.ErrRecordNotFound {
		return errno.DBOpError
	}

	if userCount != 0 && err != gorm.ErrRecordNotFound {
		return errno.PhoneHasResisted
	}


}
