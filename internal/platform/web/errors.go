package web

//ErrorResponse how we response to clients when somethings goes wrong
type ErrorResponse struct {
	Error string `json: "error"`
}

// Error is use to add web informtion to a request error
type Error struct {
	Err    error
	Status int
}

func NewRequestError(err error, status int) error {
	return &Error{Err: err, Status: status}
}

func (e *Error) Error() string {
	return e.Err.Error()
}
