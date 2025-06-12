package service

import (
	"rakuten_backend/internal/api/response/webres"
	"rakuten_backend/internal/dao"
)

func GetAdminUserOperationLogList(id int, username string, page, pageSize, operationType int, operationIp string, startTime, endTime string, isAgent int) (map[string]interface{}, error) {
	adminUserOperationLogList, count, err := dao.GetAdminUserOperationLogList(id, username, page, pageSize, operationType, operationIp, startTime, endTime, isAgent)
	if err != nil {
		return nil, err
	}
	for _, operLog := range adminUserOperationLogList {
		operLog.OperationTypeName = webres.OperationTypeMap[operLog.OperationType]
	}
	return webres.DataRsp(adminUserOperationLogList, count, page, pageSize), nil
}
