package response

type Res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (res *Res) BuildData(data any) *Res {
	res.Data = data
	return res
}

const (
	successCode         = 100000
	errCodeUnknown      = 100001
	errCodeJwtToken     = 100002
	errCodeNotFound     = 201000
	errCodeMissingParam = 202000
)

var (
	Success         = &Res{Code: successCode, Message: "Success"}
	ErrUnknown      = &Res{Code: errCodeUnknown, Message: "Unknown"}
	ErrNotFound     = &Res{Code: errCodeNotFound, Message: "Not found"}
	ErrMissingParam = &Res{Code: errCodeMissingParam, Message: "Missing Param"}
	ErrJwtToken     = &Res{Code: errCodeJwtToken, Message: "JWT Token Error"}
)
