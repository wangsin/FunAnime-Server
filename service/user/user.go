package serviceUser

import (
	"github.com/jinzhu/gorm"
	"github.com/mervick/aes-everywhere/go/aes256"
	"github.com/spf13/viper"
	"math/rand"
	"sinblog.cn/FunAnime-Server/cache"
	"sinblog.cn/FunAnime-Server/model"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
	serviceCommon "sinblog.cn/FunAnime-Server/service/common"
	"sinblog.cn/FunAnime-Server/util/consts"
	"sinblog.cn/FunAnime-Server/util/errno"
	"sinblog.cn/FunAnime-Server/util/random"
	"strconv"
	"time"
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

	flag, err := checkSmsCodeSuccess(userRequest.SmsCode)
	if err != nil {
		return errno.RedisOpError
	}

	if !flag {
		return errno.SmsCodeNotRight
	}

	_, err = model.CreateUserWithInstance(&model.User{
		Username:   "test",
		Nickname:   "test",
		Password:   aes256.Encrypt("", viper.GetString("secret_key.password_key")),
		Phone:      userRequest.Phone,
		Sex:        model.NotCommit,
		Level:      0,
		Status:     model.UserNotActive,
		Birthday:   consts.ZeroTime,
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
	})
	if err != nil {
		return errno.DBOpError
	}

	return errno.Success
}

func checkSmsCodeSuccess(smsCode string) (bool, error) {
	return true, nil
}

func SendSmsCode(phone string) error {
	smsCode := random.GenValidateCode()

	randTime := rand.Intn(3)
	minute := 15
	expireTime := time.Minute * time.Duration(minute) + time.Second * time.Duration(randTime)

	err := cache.SetSmsCode(phone, smsCode, expireTime)
	if err != nil {
		return err
	}

	// 发送短信
	err = serviceCommon.SendSms(phone, smsCode, strconv.Itoa(minute))
	if err != nil {
		return err
	}

	return nil
}