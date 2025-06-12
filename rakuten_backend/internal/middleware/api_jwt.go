package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"rakuten_backend/internal/api/auth"
	"rakuten_backend/internal/context"
	"strings"
)

// Claims 自定义的 JWT Claims
type ApiClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var (
	ApiRouterFilterMap = map[string]interface{}{
		"/api/v1/login":    struct{}{},
		"/api/v1/register": struct{}{},
	}
)

func ApiAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := context.Context{Context: ctx, Trace: ctx.GetHeader("X-Request-Id")}
		preURL := c.Request.URL.Path
		_, exist := ApiRouterFilterMap[preURL]
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

		claims, err := auth.ParseApiToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
