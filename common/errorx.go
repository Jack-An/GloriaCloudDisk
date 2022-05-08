package common

type CodeError struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

func NewCodeError(code int, err string) error {
	return &CodeError{Code: code, Err: err}
}

func NewDefaultError(err string) error {
	return NewCodeError(UNKNOWN, err)
}

func (e *CodeError) Error() string {
	return e.Err
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Err:  e.Err,
	}
}
