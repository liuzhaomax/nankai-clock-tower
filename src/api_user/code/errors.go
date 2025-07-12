package code

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Error(code int, msg string) ApiError {
	return ApiError{
		Code:    code,
		Message: msg,
	}
}

func (code ApiError) Error() string {
	return code.Message
}

var (
	InternalErr       = Error(1000, "internal error")
	RequestParsingErr = Error(1001, "请求体解析错误")
	DecodingErr       = Error(1002, "解密异常")
	EncodingErr       = Error(1003, "加密异常")
	TokenErr          = Error(1004, "Token签发错误")
	FileParsingErr    = Error(1005, "文件解析错误")

	DBFailed          = Error(2000, "数据库错误")
	UserExisted       = Error(2001, "用户已存在")
	UserNotExisted    = Error(2002, "用户不存在")
	UserUpdateFailed  = Error(2003, "更新用户信息失败")
	GroupExisted      = Error(2004, "群组已存在")
	GroupNotExisted   = Error(2005, "群组不存在")
	GroupUpdateFailed = Error(2006, "更新群组信息失败")
	UserGroupExisted  = Error(2007, "用户与群组关联已存在")

	DownstreamFailed = Error(10000, "下游服务宕机")
	WechatFailed     = Error(10001, "微信宕机")
)
