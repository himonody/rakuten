package main

import (
	"fmt"
	"rakuten_backend/config"
	"rakuten_backend/internal/router"
	"rakuten_backend/pkg/utils"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	ip, err := config.InitIp()
	if err != nil {
		panic(err)
	}
	fmt.Println(utils.GetAddressByIp("111.67.96.227"))
	defer ip.Close()
	//router := gin.Default()
	//router.GET("/ping", func(c *gin.Context) {
	//	fmt.Println(c.ClientIP())
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//router.Run() // 监听并在 0.0.0.0:8080 上启动服务
	err = router.NewRouter().Run(":8080")
	if err != nil {
		panic("failed to start server")
	}
}
