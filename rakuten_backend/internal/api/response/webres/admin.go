package webres

import "time"

var (
	OperationTypeMap = map[int]string{
		0:  "创建管理员",
		1:  "编辑管理员",
		2:  "删除管理员",
		3:  "创建会员",
		4:  "编辑会员余额",
		5:  "创建代理",
		6:  "<UNK>",
		7:  "<UNK>",
		8:  "<UNK>",
		9:  "<UNK>",
		10: "<UNK>",
		11: "<UNK>",
		12: "<UNK>",
		13: "<UNK>",
		14: "<UNK>",
	}
)

type AdminUser struct {
	Id         uint64    ` json:"id"`
	Username   string    ` json:"username"`
	Password   string    ` json:"password"`
	GoogleAuth string    `json:"google_auth"`
	AuthIP     string    `json:"auth_ip"`
	Role       uint8     `json:"role"`
	RoleName   string    `json:"role_name"`
	CreatedAt  time.Time `json:"create_at"`
	UpdatedAt  time.Time `gjson:"update_at"`
}

func DataRsp(list interface{}, count int64, page, pageSize int) map[string]interface{} {
	return map[string]interface{}{
		"list":     list,
		"count":    count,
		"page":     page,
		"pageSize": pageSize,
	}
}

type AdminUserOperationLog struct {
	ID                uint64    `gorm:"column:id" json:"id"`             // 操作id
	AdminId           uint64    `gorm:"column:admin_id" json:"admin_id"` // 管理员ID
	Username          string    `gorm:"column:username" json:"username"`
	OperationType     int       `gorm:"column:operation_type" json:"operation_type"` //操作类型
	OperationTypeName string    `gorm:"-" json:"operation_type_name"`
	OperationContent  string    `gorm:"column:operation_content" json:"operation_content"`
	OperationIp       string    `gorm:"column:operation_ip" json:"operation_ip"` //操作ip
	CreatedAt         time.Time `gorm:"column:create_at" json:"create_at"`       // 创建时间
	UpdatedAt         time.Time `gorm:"column:update_at" json:"update_at"`       // 更新时间
}
