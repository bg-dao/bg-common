package errno

var (
	OK                    = &Errno{Code: 0, Msg: "Success"}
	ParamError            = &Errno{Code: 1, Msg: "Invalid request parameters"}
	AuthError             = &Errno{Code: 2, Msg: "Login session has expired"}
	AccountOrPwdError     = &Errno{Code: 3, Msg: "Incorrect account or password"}
	HandleError           = &Errno{Code: 4, Msg: "Server is busy, please try again later"}
	ChatContextBizIdError = &Errno{Code: 5, Msg: "Chat context bizId has expired"}
	DataNotExistError     = &Errno{Code: 6, Msg: "Data does not exist"}
	OtherError            = &Errno{Code: 7, Msg: "Other error"}
)

/*var (
	OK                = &Errno{Code: 0, Msg: "成功"}
	ParamError        = &Errno{Code: 1, Msg: "请求参数错误"}
	AuthError         = &Errno{Code: 2, Msg: "登陆状态已失效"}
	AccountOrPwdError = &Errno{Code: 3, Msg: "账户或密码错误"}
	HandleError           = &Errno{Code: 4, Msg: "服务器繁忙，请稍后再试"}
	ChatContextBizIdError = &Errno{Code: 5, Msg: "对话上下文业务ID已失效"}
	DataNotExistError     = &Errno{Code: 6, Msg: "数据不存在"}
	OtherError            = &Errno{Code: 7, Msg: "其他错误"}
)
*/
