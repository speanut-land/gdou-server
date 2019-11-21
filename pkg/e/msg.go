package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_EXIST_USER:           "已存在该用户名称",
	ERROR_EXIST_USER_FAIL:      "获取已存在用户失败",
	ERROR_NOT_EXIST_USER:       "该用户不存在",
	ERROR_ADD_USER_FAIL:        "新增用户失败",
	ERROR_TELEPHONE_USED:       "该手机号已使用",
	ERROR_TELEPHONE_FORMAT:     "请输入正确的手机号",
	ERROR_LOGIN_FAIL:           "用户名或密码错误",
	ERROR_TELEPHONE_UNREGISTER: "手机号未注册",

	ERROR_CODE: "验证码错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
