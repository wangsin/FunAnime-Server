package serviceCommon

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/spf13/viper"
	"sinblog.cn/FunAnime-Server/util/logger"
	"sinblog.cn/FunAnime-Server/util/random"
	"strconv"
	"time"
)

func generateHmacSHA1(secretToken, payloadBody string) []byte {
	mac := hmac.New(sha1.New, []byte(secretToken))
	sha1.New()
	mac.Write([]byte(payloadBody))
	return mac.Sum(nil)
}

func GetVideoUploadSign() string {
	secretId := viper.GetString("tencent_api.secret_id")
	secretKey := viper.GetString("tencent_api.secret_key")
	timestamp := time.Now().Unix()
	expireTime := timestamp + 86400
	timestampStr := strconv.FormatInt(timestamp, 10)
	expireTimeStr := strconv.FormatInt(expireTime, 10)

	randStr := random.GenValidateCode()
	original := "secretId=" + secretId + "&currentTimeStamp=" + timestampStr + "&expireTime=" + expireTimeStr + "&random=" + randStr
	signature := generateHmacSHA1(secretKey, original)
	signature = append(signature, []byte(original)...)
	signatureB64 := base64.StdEncoding.EncodeToString(signature)
	logger.Info("get_upload_sign", logger.Fields{"sign": signatureB64})
	return signatureB64
}
