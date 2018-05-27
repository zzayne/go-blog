package model

type errorCode struct {
	SUCCESS      int
	ERROR        int
	NotFound     int
	LoginError   int
	LoginTimeout int
	InActive     int
	Unauthorized int
}

// ErrorCode 错误码
var ErrorCode = errorCode{
	SUCCESS:      0,
	ERROR:        1,
	NotFound:     404,
	LoginError:   1000, //用户名或密码错误
	LoginTimeout: 1001, //登录超时
	InActive:     1002, //未激活账号
	Unauthorized: 1003, //无权限
}
