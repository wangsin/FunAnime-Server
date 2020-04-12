package cache

import (
	"errors"
	"fmt"
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
