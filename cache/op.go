package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
	"sinblog.cn/FunAnime-Server/util/logger"
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

	remainDuration, err := RedisClient.TTL(key).Result()
	if err != nil {
		// 没有找到KEY 直接Set
		setErr := RedisClient.Set(key, smsCode, expireTime).Err()
		if setErr != nil {
			logger.Error("set_redis_failed_at_direct_set", logger.Fields{"err": setErr, "smsCode": smsCode, "key": key})
			return setErr
		}
		return nil
	} else {
		// 找到KEY了 判断时间 如果差值小于30s则禁止Set
		if expireTime - remainDuration*time.Second < 30 {
			return errors.New("to_quick")
		}

		// 否则set新的
		setErr := RedisClient.Set(key, smsCode, expireTime).Err()
		if setErr != nil {
			logger.Error("set_redis_failed_at_otherwise_set", logger.Fields{"err": setErr, "smsCode": smsCode, "key": key})
			return setErr
		}
		return nil
	}
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