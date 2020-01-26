package user

import (
	"github.com/jinzhu/gorm"
	"sinblog.cn/FunAnime-Server/model"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
)

func RegisterUser(userRequest user.RegisterRequestInfo) error {
	_, userCount, err := model.QueryUserWithWhereMap(
		map[string]interface{}{
			"phone": userRequest.Phone,
		},
		map[string]interface{}{
			"status <> ?": model.UserDeleted,
		},
	)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if userCount != 0 && err != gorm.ErrRecordNotFound {

	}

}
