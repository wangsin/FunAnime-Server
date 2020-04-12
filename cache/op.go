package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
	"time"
)

func SetSmsCode(phone string, smsType int, smsCode string, expireTime time.Duration) error {
	if RedisClient == nil {
		return errors.New("get_redis_db_op_failed")
	}

	key := genRedisKey(phone, smsType)
	if key == "" {
		return errors.New("params_error")
	}

	return RedisClient.Set(key, smsCode, expireTime).Err()
}

func GetSmsCode(phone string, smsType int) (string, error) {
	if RedisClient == nil {
		return "", errors.New("get_redis_db_op_failed")
	}

	key := genRedisKey(phone, smsType)
	if key == "" {
		return "", errors.New("params_error")
	}

	strCmd := RedisClient.Get(key)
	if strCmd == nil || strCmd.Err() != nil {
		return "", errors.New("redis_op_failed_at_getSmsCode")
	}
	return strCmd.Val(), nil
}

func genRedisKey(phone string, smsType int) string {
	switch smsType {
	case user.Register:
		return fmt.Sprintf(RedisRegisterSmsCodeKey, phone)
	case user.Login:
		return fmt.Sprintf(RedisLoginSmsCodeKey, phone)
	default:
		return ""
	}
}

func SetUserLogin(expTime time.Duration, userInfo *token.UserInfo) error {
	if RedisClient == nil {
		return errors.New("get_redis_db_op_failed")
	}

	return RedisClient.Set(fmt.Sprintf(RedisUserLoginKey, userInfo.UserId), userInfo.ToJson(), expTime).Err()
}

func GetUserLogin(userId int64) (*token.UserInfo, error) {
	if RedisClient == nil {
		return nil, errors.New("get_redis_db_op_failed")
	}

	strCmd := RedisClient.Get(fmt.Sprintf(RedisUserLoginKey, userId))
	if strCmd == nil || strCmd.Err() != nil {
		return nil, errors.New("redis_op_failed_at_GetUserLogin")
	}

	userInfo := new(token.UserInfo)

	err := json.Unmarshal([]byte(strCmd.Val()), userInfo)
	if err != nil {
		return nil, errors.New("json_unmarshal_failed")
	}

	return userInfo, nil
}

func DelUserLogin(userId int64) error {
	if RedisClient == nil {
		return errors.New("get_redis_db_op_failed")
	}

	return RedisClient.Del(fmt.Sprintf(RedisUserLoginKey, userId)).Err()
}