package serviceUser

import (
	"github.com/jinzhu/gorm"
	"github.com/mervick/aes-everywhere/go/aes256"
	"github.com/spf13/viper"
	"sinblog.cn/FunAnime-Server/cache"
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/model"
	reqUser "sinblog.cn/FunAnime-Server/serializable/request/user"
	serviceCommon "sinblog.cn/FunAnime-Server/service/common"
	"sinblog.cn/FunAnime-Server/util/consts"
	"sinblog.cn/FunAnime-Server/util/errno"
	"sinblog.cn/FunAnime-Server/util/logger"
	"sinblog.cn/FunAnime-Server/util/random"
	"strconv"
	"time"
)

func RegisterUser(userRequest *reqUser.RegisterRequestInfo) int64 {
	result, userCount, err := model.QueryUserWithWhereMap(
		map[string]interface{}{
			"phone": userRequest.Phone,
		},
		map[string]interface{}{
			"status <> ?": model.UserDeleted,
		},
	)

	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("db_op_failed_at_QueryUserWithWhereMap", logger.Fields{"err": err.Error(), "request": userRequest})
		return errno.DBOpError
	}

	if len(result) > 0 || userCount > 0 {
		logger.Error("RegisterUser_PhoneHasResisted", logger.Fields{"request": userRequest})
		return errno.PhoneHasResisted
	}

	flag, err := checkSmsCodeSuccess(userRequest.Phone, userRequest.SmsCode, reqUser.Register)
	if err != nil {
		logger.Error("RegisterUser_SmsCodeNotSend", logger.Fields{"err": err.Error(), "request": userRequest})
		return errno.SmsCodeNotSend
	}

	if !flag {
		logger.Error("RegisterUser_SmsCodeNotRight", logger.Fields{"request": userRequest})
		return errno.SmsCodeNotRight
	}

	password := userRequest.Password
	if len(userRequest.Password) <= 0 {
		password = aes256.Encrypt(random.GenRandomPassword(), viper.GetString("secret_key.password_key"))
	}

	_, err = model.CreateUserWithInstance(&model.User{
		Username:   random.GenEncryptUserName(userRequest.Phone),
		Nickname:   random.GenEncryptUserName(userRequest.Phone),
		Mail:       userRequest.Mail,
		Password:   password,
		Phone:      userRequest.Phone,
		Sex:        model.NotCommit,
		Status:     model.UserNotActive,
		Birthday:   consts.ZeroTime,
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
	})

	if err != nil {
		logger.Error("RegisterUser_CreateUserWithInstance", logger.Fields{"err": err.Error(), "request": userRequest})
		return errno.DBOpError
	}

	return errno.Success
}

func checkSmsCodeSuccess(phone, smsCode string, smsType int) (bool, error) {
	sCode, err := cache.GetSmsCode(phone, smsType)
	if err != nil || sCode == "" {
		return false, err
	}
	if sCode != smsCode {
		return false, nil
	}
	return true, nil
}

func SendSmsCode(request *reqUser.SendSmsRequest) error {
	smsCode := random.GenValidateCode()

	minute := 300
	expireTime := time.Second * time.Duration(minute)

	err := cache.SetSmsCode(request.Phone, request.Type, smsCode, expireTime)
	if err != nil {
		logger.Error("set_sms_code_failed", logger.Fields{"err": err, "request": request})
		return err
	}

	// 发送短信
	err = serviceCommon.SendSms(request.Phone, smsCode, strconv.Itoa(minute/60))
	if err != nil {
		logger.Error("send_sms_error", logger.Fields{"err": err, "request": request})
		return err
	}

	return nil
}

func checkPasswordRight(requestPassword, dbPassword string) bool {
	key := viper.GetString("secret_key.password_key")
	return requestPassword == aes256.Decrypt(dbPassword, key)
}

func LoginUser(userRequest *reqUser.LoginRequestInfo) (string, *model.User, int64) {
	userList, userCount, err := model.QueryUserWithWhereMap(
		map[string]interface{}{
			"phone": userRequest.Phone,
		},
		map[string]interface{}{
			"status <> ?": model.UserDeleted,
		},
	)

	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("db_op_failed", logger.Fields{"err": err, "request": userRequest})
		return "", nil, errno.DBOpError
	}

	if err == gorm.ErrRecordNotFound || userCount == 0 || len(userList) <= 0 {
		logger.Error("phone_not_exist", logger.Fields{"err": err, "request": userRequest})
		return "", nil, errno.PhoneNotExistence
	}

	userInfo := userList[0]
	flag := false
	if userRequest.Password != "" {
		flag = checkPasswordRight(userRequest.Password, userInfo.Password)
	} else if userRequest.SmsCode != "" {
		flag, err = checkSmsCodeSuccess(userRequest.Phone, userRequest.SmsCode, reqUser.Login)
		if err != nil {
			logger.Error("check_sms_code_failed", logger.Fields{"err": err, "request": userRequest})
			return "", nil, errno.SmsCodeNotSend
		}
	}

	if !flag {
		logger.Error("login_info_not_fit", logger.Fields{"err": err, "request": userRequest})
		return "", nil, errno.LoginInfoFailed
	}

	tokenUserInfo := &token.UserInfo{
		UserId:   userInfo.Id,
		Phone:    userInfo.Phone,
		Nickname: userInfo.Nickname,
		Username: userInfo.Username,
	}

	tokenUserInfo.ExpiresAt = time.Now().AddDate(0, 0, 15).Unix()

	tToken, err := token.NewJWT().CreateToken(tokenUserInfo)
	if err != nil {
		logger.Error("token_generate_failed", logger.Fields{"err": err, "request": userRequest})
		return "", nil, errno.TokenInvalid
	}

	if err := cache.SetUserLogin(time.Hour*24*10, tokenUserInfo); err != nil {
		logger.Error("set_redis_op_failed", logger.Fields{"err": err, "request": userRequest})
		return "", nil, errno.RedisOpError
	}

	return tToken, userInfo, errno.Success
}

func GetUserInfo(userInfo *reqUser.BasicUser) (*model.User, int64) {
	dbUserInfo, err := model.QueryUserWithId(userInfo.UserInfo.UserId)
	if err != nil || dbUserInfo == nil {
		return nil, errno.DBOpError
	}

	return dbUserInfo, errno.Success
}

func Logout(userInfo *reqUser.BasicUser) error {
	if err := cache.DelUserLogin(userInfo.UserInfo.UserId); err != nil {
		return err
	}

	return nil
}
