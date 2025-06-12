package request

type AdminUser struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	AuthIp   *string `json:"auth_ip"`
	Status   *int    `json:"status"`
	Role     *int    `json:"role"`
}
