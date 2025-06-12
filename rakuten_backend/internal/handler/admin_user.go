package handler

import (
	"rakuten_backend/internal/api/request"
	"rakuten_backend/internal/api/response/code"
	"rakuten_backend/internal/api/response/web"
	"rakuten_backend/internal/context"
	"rakuten_backend/internal/service"
	"rakuten_backend/pkg/utils"
	"strconv"
	"strings"
)

func CreateAdmin(c *context.Context) {
	admin, err := c.GetAdminClaims()
	if admin == nil || err != nil {
		c.Errorf("创建管理员账号失败1: %v", err)
		web.Fail(c, code.PermissionDenied)
		return
	}

	if admin.Role != 1 {
		c.Errorf("创建管理员账号失败2: %v", err)
		web.Fail(c, code.PermissionDenied)
		return
	}

	req := new(request.AdminAuth)
	if err = c.ShouldBind(req); err != nil {
		c.Errorf("创建管理员账号失败3: %v", err)
		web.Fail(c, code.InvalidParam)
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	if req.Username == "" || req.Password == "" {
		web.Fail(c, code.UserPasswordError)
	}

	err = service.CreateAdmin(req)
	if err != nil {
		c.Errorf("创建管理员账号失败4: %v", err)
		web.Fail(c, code.InternalError)
		return
	}

	web.Success(c, "")
	return
}
func AdminList(c *context.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	username := c.Query("username")
	role, _ := strconv.Atoi(c.DefaultQuery("role", "-1"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	dataRes, err := service.GetAdminList(username, id, page, pageSize, role, 0)
	if err != nil {
		c.Errorf("查询管理员列表失败1: %v", err)
		web.Fail(c, code.InternalError)
		return
	}

	web.Success(c, dataRes)
	return
}

func EditAdminUser(c *context.Context) {
	admin, err := c.GetAdminClaims()
	if admin == nil || err != nil {
		c.Errorf("编辑管理员账号失败1: %v", err)
		web.Fail(c, code.PermissionDenied)
		return
	}

	if admin.Role != 1 {
		c.Errorf("编辑管理员账号失败2: %v", err)
		web.Fail(c, code.PermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Errorf("编辑管理员账号失败3: %v", err)
		web.Fail(c, code.InvalidParam)
		return
	}
	req := new(request.AdminUser)
	if err = c.ShouldBind(req); err != nil {
		c.Errorf("编辑管理员账号失败4: %v", err)
		web.Fail(c, code.InvalidParam)
		return
	}
	adminUserMap := make(map[string]interface{})
	if req.Username != "" {
		adminUserMap["username"] = req.Username
	}
	if req.Password != "" {
		adminUserMap["password"] = utils.EncryptionMD5(req.Password)
	}
	if *req.Role >= 0 {
		adminUserMap["role"] = *req.Role
	}
	if *req.Status >= 0 {
		adminUserMap["status"] = *req.Status
	}
	if *req.AuthIp != "" {
		adminUserMap["auth_ip"] = *req.AuthIp
	}
	err = service.EditAdminUser(id, adminUserMap)
	if err != nil {
		c.Errorf("编辑管理员账号失败5: %v", err)
		web.Fail(c, code.InternalError)
		return
	}

	web.Success(c, "")
	return
}

func DeleteAdminUser(c *context.Context) {
	admin, err := c.GetAdminClaims()
	if admin == nil || err != nil {
		c.Errorf("删除管理员账号失败1: %v", err)
		web.Fail(c, code.PermissionDenied)
		return
	}

	if admin.Role != 1 {
		c.Errorf("删除管理员账号失败2: %v", err)
		web.Fail(c, code.PermissionDenied)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Errorf("删除管理员账号失败3: %v", err)
		web.Fail(c, code.InvalidParam)
		return
	}
	err = service.DeleteAdminUser(id)
	if err != nil {
		c.Errorf("删除管理员账号失败4: %v", err)
		web.Fail(c, code.InternalError)
		return
	}
	web.Success(c, "")
	return
}
