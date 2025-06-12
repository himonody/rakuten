package cmd

import (
	"rakuten_backend/config"
	"rakuten_backend/internal/db"
	"rakuten_backend/pkg/log"
)

func Init() error {

	if err := config.InitConfig(); err != nil {
		return err
	}
	if err := log.InitLog(); err != nil {
		return err
	}
	if err := db.InitMySQL(); err != nil {
		return err
	}
	return nil
}
