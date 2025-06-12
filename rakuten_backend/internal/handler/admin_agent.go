package handler

import (
	"rakuten_backend/internal/api/request"
	"rakuten_backend/internal/api/response/code"
	"rakuten_backend/internal/api/response/web"
	"rakuten_backend/internal/context"
	"rakuten_backend/internal/service"
	"strconv"
	"strings"
)

func CreateAgent(c *context.Context) {
	admin, err := c.GetAdminClaims()
	if admin == nil || err != nil {
		c.Errorf("创建代理账号失败1: %v", err)
		web.Fail(c, code.PermissionDenied)
		return
	}

	if admin.Role > 1 {
		c.Errorf("创建代理账号失败2: %v", err)
		web.Fail(c, code.PermissionDenied)
		return
	}

	req := new(request.AdminAuth)
	if err = c.ShouldBind(req); err != nil {
		c.Errorf("创建代理账号失败3: %v", err)
		web.Fail(c, code.InvalidParam)
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	if req.Username == "" || req.Password == "" {
		web.Fail(c, code.UserPasswordError)
	}

	err = service.CreateAgent(req)
	if err != nil {
		c.Errorf("创建代理账号失败4: %v", err)
		web.Fail(c, code.InternalError)
		return
	}

	web.Success(c, "")
	return
}

func AgentList(c *context.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	username := c.Query("username")
	role, _ := strconv.Atoi(c.DefaultQuery("role", "-1"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	dataRes, err := service.GetAdminList(username, id, page, pageSize, role, 1)
	if err != nil {
		c.Errorf("查询管理员列表失败1: %v", err)
		web.Fail(c, code.InternalError)
		return
	}

	web.Success(c, dataRes)
	return
}
