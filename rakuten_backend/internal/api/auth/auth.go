package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"

	"rakuten_backend/config"

	"time"
)

// Claims 自定义的 JWT Claims
type AdminClaims struct {
	AdminID   uint64 `json:"admin_id"`
	AdminName string `json:"admin_name"`
	Role      uint8  `json:"role"`
	IsAgent   uint8  `json:"is_agent"`
	jwt.RegisteredClaims
}

// GenerateAdminToken  生成 JWT token
func GenerateAdminToken(userID uint64, username string, role, isAgent uint8) (string, error) {
	claims := AdminClaims{
		AdminID:   userID,
		AdminName: username,
		Role:      role,
		IsAgent:   isAgent,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.GetJWT().Expire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.GetJWT().Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetJWT().Secret))
}

// ParseAdminToken 解析 JWT token
func ParseAdminToken(tokenString string) (*AdminClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWT().Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// Claims 自定义的 JWT Claims
type ApiClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`

	jwt.RegisteredClaims
}

// GenerateAdminToken  生成 JWT token
func GenerateApiToken(userID uint64, username string) (string, error) {
	claims := ApiClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.GetJWT().Expire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.GetJWT().Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetJWT().Secret))
}

// ParseApiToken 解析 JWT token
func ParseApiToken(tokenString string) (*ApiClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &ApiClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWT().Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*ApiClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
