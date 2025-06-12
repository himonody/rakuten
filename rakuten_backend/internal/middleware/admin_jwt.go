package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rakuten_backend/internal/api/auth"
	"rakuten_backend/internal/context"
	"strings"
)

var (
	AdminRouterFilterMap = map[string]interface{}{
		"/admin/v1/login": struct{}{},
	}
)

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := context.Context{Context: ctx, Trace: ctx.GetHeader("X-Request-Id")}
		preURL := c.Request.URL.Path
		_, exist := AdminRouterFilterMap[preURL]
		if exist {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		claims, err := auth.ParseAdminToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("admin_id", claims.AdminID)
		c.Set("admin_name", claims.AdminName)
		c.Set("role", claims.Role)
		c.Set("is_agent", claims.IsAgent)

		c.Next()
	}
}
