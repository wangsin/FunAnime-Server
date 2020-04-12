package random

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	ValidateCodeLength = 6
)

func GenValidateCode() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < ValidateCodeLength; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func GenEncryptUserName(phone string) string {
	baseUser := "%s_USER_PHONE"
	username := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(baseUser, phone)))
	if len(username) >= 10 {
		return username[:10]
	} else {
		return username
	}
}

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func MD5(text string) string {
	newMd5 := md5.New()
	newMd5.Write([]byte(text))
	return hex.EncodeToString(newMd5.Sum(nil))
}

func GenRandomPassword() string {
	return MD5(GetRandomString(20))
}
