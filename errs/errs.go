package errs

type Errs struct {
	ErrCode string
	Message string
}

func (e *Errs) Error() string {
	return e.Message
}

func NewComplexErrs(errCode string,message string) error {
	return &Errs{ErrCode:errCode,Message:message}
}




