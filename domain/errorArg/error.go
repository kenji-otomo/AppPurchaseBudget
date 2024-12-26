package errorArg

type Error struct {
	ErrMsg string `json:"error_message"`
}

func NewError(err string) *Error {
	return &Error{
		ErrMsg: err,
	}
}
