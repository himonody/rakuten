package handler

import (
	"rakuten_backend/internal/api/auth"
	"rakuten_backend/internal/api/request"
	"rakuten_backend/internal/api/response/code"
	"rakuten_backend/internal/api/response/web"
	"rakuten_backend/internal/context"
	"rakuten_backend/internal/service"
)

func AdminLogin(c *context.Context) {
	req := new(request.AdminAuth)
	if err := c.ShouldBind(req); err != nil {
		c.Errorf("管理员账号登录失败1: %v", err)
		web.Fail(c, code.InvalidParam)
	}
	user, err := service.AdminLogin(req)
	if err != nil {
		c.Errorf("管理员账号登录失败2: %v", err)
		web.Fail(c, code.UserNoExist)
		return
	}
	if !service.ValidatePassword(req.Password, user.Password) {
		c.Errorf("管理员账号登录失败3: %v 用户密码错误 ", req.Password)
		web.Fail(c, code.UserPasswordError)
		return
	}
	if !service.ValidateAuthIP(c.ClientIP(), user.AuthIP) {
		c.Errorf("管理员账号登录失败4: %v 该 IP 地址被限制访问 ", c.ClientIP())
		web.Fail(c, code.IPNotAllowed)
		return
	}
	if !service.ValidateGoogleCode(req.Code, user.GoogleAuth) {
		c.Errorf("管理员账号登录失败5: %v 谷歌验证码错误 ", req.Code)
		web.Fail(c, code.GoogleCodeError)
		return
	}

	token, err := auth.GenerateAdminToken(user.ID, user.Username, user.Role, user.Role)
	if err != nil {
		c.Errorf("管理员账号登录失败6: %v", err)
		web.Fail(c, code.InternalError)
		return
	}
	web.Success(c, token)
	return
}
