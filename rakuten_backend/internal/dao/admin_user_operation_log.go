package dao

import (
	"rakuten_backend/internal/api/response/webres"
	"rakuten_backend/internal/db"
)

func GetAdminUserOperationLogList(id int, username string, page, pageSize, operationType int, operationIp string, startTime, endTime string, isAgent int) (
	adminUserOperationLogList []*webres.AdminUserOperationLog, count int64, err error) {
	adminUserOperationLogList = make([]*webres.AdminUserOperationLog, 0)
	res := db.MySQL.Model("admin_user_operation_log as ol")
	res.Select("ol.id,ol.admin_id,au.username,ol.operation_type,ol.operation_ip,ol.update_at")
	res.Joins("left join admin_user as au on au.admin_id = ol.id")
	res.Where("au.is_agent = ?", isAgent)
	if id > 0 {
		res.Where("ol.id = ?", id)
	}
	if username != "" {
		res.Where("au.username = ?", username)
	}
	if operationType > -1 {
		res.Where("ol.operation_type = ?", operationType)
	}
	if operationIp != "" {
		res.Where("ol.operation_ip = ?", operationIp)
	}
	if startTime != "" {
		res.Where("ol.create_at >= ?", startTime)
	}
	if endTime != "" {
		res.Where("ol.update_at <= ?", endTime)
	}
	if err = res.Count(&count).Error; err != nil || count < 1 {

		return adminUserOperationLogList, 0, err
	}
	res.Order("ol.id desc")
	if err = res.Offset((page - 1) * pageSize).Limit(pageSize).Find(&adminUserOperationLogList).Error; err != nil {
		return adminUserOperationLogList, 0, err
	}
	return adminUserOperationLogList, count, nil

}
