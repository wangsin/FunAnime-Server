package cache

import (
	"errors"
	"fmt"
	"time"
)

func SetSmsCode(phone string, smsCode string, expireTime time.Duration) error {
	if RedisClient == nil {
		return errors.New("get_redis_db_op_failed")
	}

	return RedisClient.Set(fmt.Sprintf(RedisSmsCodeKey, phone), smsCode, expireTime).Err()
}
