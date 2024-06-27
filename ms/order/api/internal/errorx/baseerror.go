package errorx

// 自定义的错误

const (
	defaultErrCode = 1001
	RPCErrCode     = 1002
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) CodeError {
	return CodeError{
		Code: code,
		Msg:  msg,
	}
}

func NewDefaultCodeError(msg string) error {
	return NewCodeError(defaultErrCode, msg)
}

// CodeError实现error的接口
func (e CodeError) Error() string {
	return e.Msg
}

// Data 返回自定义类型的错误响应时使用
func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

