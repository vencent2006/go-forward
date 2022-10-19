package controller

type ResCode int

const (
	// CodeSuccess todo 用立即数比较好，客户端是根据数值进行判断，变化了影响会很大; 只能顺序新增，不能插入和删除
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
	CodeCommunityNotExist
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:           "成功",
	CodeInvalidParam:      "请求参数错误",
	CodeUserExist:         "用户名已存在",
	CodeUserNotExist:      "用户名不存在",
	CodeInvalidPassword:   "用户名或密码错误",
	CodeServerBusy:        "服务繁忙",
	CodeNeedLogin:         "需要登录",
	CodeInvalidToken:      "无效的token",
	CodeCommunityNotExist: "社区不存在",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
