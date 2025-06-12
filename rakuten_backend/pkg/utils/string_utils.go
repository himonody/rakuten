package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// MD5WithSalt 加盐的 MD5 加密函数
func MD5WithSalt(password string) string {
	salt := "MyFixedSalt@" + password
	salt = ReverseString(salt)
	data := []byte(password + salt)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// ReverseString 反转字符串，支持 Unicode 字符（如中文、emoji），对空字符串自动处理
func ReverseString(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
func EncryptionMD5(str string) string {
	return ReverseString(MD5WithSalt(ReverseString(MD5WithSalt(str))))
}

// generateAlphaNumeric 随机数
func generateAlphaNumeric(length int) string {
	if length <= 0 {
		return ""
	}

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const charset = letters + "0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)

	// 第一个字符必须是字母
	result[0] = letters[r.Intn(len(letters))]

	// 剩余字符可为字母或数字
	for i := 1; i < length; i++ {
		result[i] = charset[r.Intn(len(charset))]
	}

	return string(result)
}
func GenerateRandomString(length int) string {
	return time.Now().Format("20060102150405") + generateAlphaNumeric(length)
}

// GetActivityId 活动ID
func GetActivityId() string {
	return "AC" + GenerateRandomString(7)
}

// GetTopUpId 充值ID
func GetTopUpId() string {
	return "TU" + GenerateRandomString(7)
}

// GetPayOutId 提款ID
func GetPayOutId() string {
	return "PO" + GenerateRandomString(7)
}

// GetRebateId 返利ID
func GetRebateId() string {
	return "RE" + GenerateRandomString(7)
}

// GetSnapUpId 提款ID
func GetSnapUpId() string {
	return "SU" + GenerateRandomString(7)
}

// GetGroupBuyId 团队任务ID
func GetGroupBuyId() string {
	return "GB" + GenerateRandomString(7)
}
