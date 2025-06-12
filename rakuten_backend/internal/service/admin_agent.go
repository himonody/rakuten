package service

import (
	"rakuten_backend/internal/api/request"
	"rakuten_backend/internal/dao"
	"rakuten_backend/internal/model"
	"rakuten_backend/pkg/utils"
)

func CreateAgent(req *request.AdminAuth) (err error) {
	adminUser := new(model.AdminUser)
	adminUser.Username = req.Username
	adminUser.Password = utils.EncryptionMD5(req.Password)
	adminUser.GoogleAuth, err = GenerateSimpleGoogleSecret(adminUser.Username)
	adminUser.IsAgent = 1
	if err != nil {
		return err
	}
	err = dao.InsertAdminUser(adminUser)
	if err != nil {
		return err
	}
	return nil
}
