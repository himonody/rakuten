package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rakuten_backend/config"
	"rakuten_backend/internal/api/xhttp"
	"rakuten_backend/internal/context"
	"rakuten_backend/internal/middleware"
	"rakuten_backend/pkg/log"
	"time"
)

var (
	api, admin *gin.RouterGroup
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(log.GinLogger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	if config.GlobalConfig.HTTP.Timeout > 0 {
		// if you need more fine-grained control over your routes, set the timeout in your routes, unsetting the timeout globally here.
		r.Use(middleware.Timeout(time.Second * time.Duration(config.GlobalConfig.HTTP.Timeout)))
	}

	// 健康检查
	xhttp.GET(r.Group("/"), "health", func(c *context.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "ok",
			"message": "Hello Rakuten!",
		})
		return
	})
	api = r.Group("/api/v1", middleware.AdminAuth())
	setApiRouter()
	admin = r.Group("/admin/v1", middleware.ApiAuth())
	setAdminRouter()

	return r
}
