package request

type AdminAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
}
