package web

import (
	"net/http"
	"rakuten_backend/internal/api/response/code"
	"rakuten_backend/internal/context"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func Success[T any](c *context.Context, data T) {
	c.JSON(http.StatusOK, Response[T]{
		Code:    code.Success.Code,
		Message: code.Success.Msg,
		Data:    data,
	})
}

func Fail(c *context.Context, code code.ErrCode) {
	c.JSON(http.StatusOK, Response[any]{
		Code:    code.Code,
		Message: code.Msg,
		Data:    nil,
	})
}
