package utils

import (
	"fmt"
	"rakuten_backend/config"
	"testing"
)

func TestGetActivityId(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(GetActivityId())
	}
}

func TestGetActivityId2(t *testing.T) {
	fmt.Println(EncryptionMD5("admin123"))
	_, err := config.InitIp()
	if err != nil {
		fmt.Printf("failed to init ip: %s", err)
	}
	//fmt.Println(GetAddressByIp("111.67.96.227"))
}
