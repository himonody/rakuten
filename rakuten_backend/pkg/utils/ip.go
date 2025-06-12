package utils

import (
	"fmt"
	"rakuten_backend/config"
	"strings"
)

func GetAddressByIp(ip string) string {
	region, err := config.Searcher.SearchByStr(ip)
	if err != nil {
		return ""
	}
	parts := strings.Split(region, "|")
	address := ""
	if len(parts) >= 5 {
		address = parts[0]
		if parts[1] != "0" {
			address = fmt.Sprintf("%s-%s", address, parts[1])
		}
		if parts[2] != "0" {
			address = fmt.Sprintf("%s-%s", address, parts[2])
		}
		if parts[3] != "0" {
			address = fmt.Sprintf("%s-%s", address, parts[3])
		}
		if parts[4] != "0" {
			address = fmt.Sprintf("%s-%s", address, parts[4])
		}
	}
	return address
}
