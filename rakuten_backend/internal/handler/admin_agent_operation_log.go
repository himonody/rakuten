package handler

import (
	"rakuten_backend/internal/api/response/code"
	"rakuten_backend/internal/api/response/web"
	"rakuten_backend/internal/context"
	"rakuten_backend/internal/service"
	"rakuten_backend/pkg/utils"
	"strconv"
)

func AdminAgentOperationLogList(c *context.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	username := c.Query("username")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	operationType, _ := strconv.Atoi(c.DefaultQuery("operationType", "-1"))
	operationIp := c.ClientIP()
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")

	if startTime == "" {
		startTime = utils.StartOfDay()
	}
	if endTime == "" {
		endTime = utils.EndOfDay()
	}

	dataRsp, err := service.GetAdminAgentOperationLogList(id, username, page, pageSize, operationType, operationIp, startTime, endTime, 1)
	if err != nil {
		c.Errorf("查询管理管理员操作日志失败: %v", err)
		web.Fail(c, code.InternalError)
		return
	}
	web.Success(c, dataRsp)
	return
}
