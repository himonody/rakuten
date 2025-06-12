package service

import (
	"errors"
	"github.com/pquerna/otp/totp"
	"rakuten_backend/internal/api/request"
	"rakuten_backend/internal/api/response/webres"
	"rakuten_backend/internal/dao"
	"rakuten_backend/internal/model"
	"rakuten_backend/pkg/utils"
	"strings"
)

func AdminLogin(req *request.AdminAuth) (user *model.AdminUser, err error) {
	user, err = dao.GetAdminUserByUserName(req.Username)
	if err != nil {
		return nil, err
	}

	if user == nil || user.GoogleAuth == "" {
		return nil, errors.New("管理员用户不存在")
	}

	return user, nil
}
func ValidateAuthIP(clientIp, authIp string) bool {
	return strings.Contains(authIp, clientIp)
}
func GenerateSimpleGoogleSecret(username string) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		AccountName: username,
	})
	if err != nil {
		return "", err
	}
	return key.Secret(), nil
}
func ValidateGoogleCode(code, secret string) bool {
	return totp.Validate(code, secret)
}
func ValidatePassword(password, secret string) bool {
	return utils.EncryptionMD5(password) == secret
}
func CreateAdmin(req *request.AdminAuth) (err error) {
	adminUser := new(model.AdminUser)
	adminUser.Username = req.Username
	adminUser.Password = utils.EncryptionMD5(req.Password)
	adminUser.GoogleAuth, err = GenerateSimpleGoogleSecret(adminUser.Username)
	adminUser.IsAgent = 0
	if err != nil {
		return err
	}
	err = dao.InsertAdminUser(adminUser)
	if err != nil {
		return err
	}
	return nil
}
func GetAdminList(username string, id, role, page, pageSize, isAgent int) (data map[string]interface{}, err error) {

	admins, count, err := dao.GetAdminUser(username, id, role, page, pageSize, isAgent)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return
	}
	adminList := make([]*webres.AdminUser, 0, len(admins))
	for _, adminUser := range admins {
		roleName := "超级管理员"
		if adminUser.Role == 0 {
			roleName = "普通管理员"
		}
		adminList = append(adminList, &webres.AdminUser{
			Id:         adminUser.ID,
			Username:   adminUser.Username,
			Role:       adminUser.Role,
			Password:   adminUser.Password,
			GoogleAuth: adminUser.GoogleAuth,
			CreatedAt:  adminUser.CreatedAt,
			UpdatedAt:  adminUser.UpdatedAt,
			RoleName:   roleName,
		})
	}

	return webres.DataRsp(adminList, count, page, pageSize), err
}

func EditAdminUser(id int, adminUserMap map[string]interface{}) (err error) {
	err = dao.EditAdminUser(id, adminUserMap)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAdminUser(id int) (err error) {
	err = dao.DeleteAdminUser(id)
	if err != nil {
		return err
	}
	return nil
}
