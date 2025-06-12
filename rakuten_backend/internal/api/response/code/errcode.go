package code

type ErrCode struct {
	Code int
	Msg  string
}

var (
	Success       = ErrCode{0, "success"}
	InvalidParam  = ErrCode{1001, "invalid parameter"}
	Unauthorized  = ErrCode{1002, "unauthorized"}
	InternalError = ErrCode{5000, "internal server error"}
)
var (
	IPNotAllowed      = ErrCode{1001, "Login failed: This IP address is restricted"}
	PermissionDenied  = ErrCode{1001, "Permission denied. You are not authorized to perform this action"}
	GoogleCodeError   = ErrCode{1001, "Google code error"}
	UserNoExist       = ErrCode{1001, "User no exist"}
	UserPasswordError = ErrCode{1001, "User password error"}
)
