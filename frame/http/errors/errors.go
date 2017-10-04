package errors

import (
	"net/http"
)

// Error 包含发生该错误时的返回码与返回信息
type Error struct {
	Code int    // 响应码
	Msg  string // 响应英文简述
	Text string // 响应中文详述
}

// newError 返回一种新的错误类型
func newError(code int, msg string, text string) *Error {
	return &Error{code, msg, text}
}

// Err 未注册的错误类型，直接输入错误信息，返回Error类型指针
func Err(info string) *Error {
	return &Error{403, "UnRegisteredError", info}
}

// 错误类型集合
var (
	ErrMethodNotSupport = newError(http.StatusForbidden, "MethodNotSupport", "该URL不支持当前模式的访问")
	ErrParamIsWrong     = newError(http.StatusForbidden, "ParamIsWrong", "参数错误")
	ErrAuthNotPass      = newError(http.StatusForbidden, "AuthNotPass", "账号与密码不匹配")
)
