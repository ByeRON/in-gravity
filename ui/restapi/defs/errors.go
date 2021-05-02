package defs

type ResponseError struct {
	ErrorCode string
	Message   string
}

func NewResponseError(e error) (int, ResponseError) {
	var httpStatus int
	return httpStatus, ResponseError{}
}
