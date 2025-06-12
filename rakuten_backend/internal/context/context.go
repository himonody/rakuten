package context

import (
	"errors"
	"github.com/gin-gonic/gin"
	"path"
	"rakuten_backend/internal/api/auth"
	"rakuten_backend/pkg/log"
	"runtime"
	"strings"
)

type Handler func(ctx *Context)

type Context struct {
	*gin.Context
	Trace string
}

func Background(c *gin.Context) *Context {
	return &Context{c, c.GetHeader("X-Request-Id")}
}

func (c *Context) Info(args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Info(args...)
}

func (c *Context) Infof(template string, args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Infof(template, args...)
}

func (c *Context) Warn(args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Warn(args...)
}

func (c *Context) Warnf(template string, args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Warnf(template, args...)
}

func (c *Context) Error(args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Error(args...)
}

func (c *Context) Errorf(template string, args ...interface{}) {
	log.ZapLog.Named(funcName(c.Trace)).Errorf(template, args...)
}

func funcName(trace string) string {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return path.Base(funcName) + " " + trace
}

func (c *Context) GetAdminClaims() (*auth.AdminClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {

		return nil, errors.New("token 失效")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {

		return nil, errors.New("token 解析失败")
	}

	claims, err := auth.ParseAdminToken(parts[1])
	if err != nil || claims == nil {

		return nil, err
	}
	return claims, err
}
func (c *Context) GetApiClaims() (*auth.ApiClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {

		return nil, errors.New("token 失效")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {

		return nil, errors.New("token 解析失败")
	}

	claims, err := auth.ParseApiToken(parts[1])
	if err != nil || claims == nil {

		return nil, err
	}
	return claims, err
}
