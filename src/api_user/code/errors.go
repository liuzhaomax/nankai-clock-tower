package code

import "github.com/liuzhaomax/go-maxms/internal/core"

func Error(code int, msg string) *core.ApiError {
	return &core.ApiError{
		Code:    code,
		Message: msg,
	}
}

var (
	// 内部服务器错误
	InternalErr       = Error(1000, "内部错误")
	DBFailed          = Error(1001, "数据库错误")
	PukGenErr         = Error(1100, "公钥生成失败")
	RequestParsingErr = Error(1101, "请求体解析错误")
	DecodingErr       = Error(1102, "解密异常")
	EncodingErr       = Error(1103, "加密异常")
	TokenErr          = Error(1104, "Token签发错误")
	FileParsingErr    = Error(1105, "文件解析错误")
	// 数据错误
	UserExisted       = Error(2001, "用户已存在")
	UserNotExisted    = Error(2002, "用户不存在")
	UserUpdateFailed  = Error(2003, "更新用户信息失败")
	GroupExisted      = Error(2004, "群组已存在")
	GroupNotExisted   = Error(2005, "群组不存在")
	GroupUpdateFailed = Error(2006, "更新群组信息失败")
	UserGroupExisted  = Error(2007, "用户与群组关联已存在")
	QuitGroupFailed   = Error(2008, "退组失败")
	DeleteGroupFailed = Error(2009, "删除群组失败")
	// 下游服务错误
	DownstreamFailed = Error(10000, "下游服务宕机")
	WechatFailed     = Error(10001, "微信宕机")
)
