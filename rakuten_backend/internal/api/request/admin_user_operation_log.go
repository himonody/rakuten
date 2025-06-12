package request

type AdminUserOperationLog struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	OperationType string `json:"operation_type"`
	OperationIp   string `json:"operation_ip"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
}
