package usecase

type ResultStatus struct {
	Code  int
	Error error
}

func NewResultStatus(code int, err error) *ResultStatus {
	return &ResultStatus{
		Code:  code,
		Error: err,
	}
}
