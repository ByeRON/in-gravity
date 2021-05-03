package defs

import (
	domainErrors "in-gravity/domains/errors"
	uiErrors "in-gravity/ui/restapi/errors"
	"net/http"
)

type ResponseError struct {
	ErrorCode string
	Message   string
}

var ErrorCodeHttpStatusMap = map[domainErrors.ErrorCodeInterface]int{
	uiErrors.ErrorSomething: http.StatusBadRequest,
}

func NewResponseError(e error) (int, ResponseError) {
	var httpStatus int
	errorCode, err := uiErrors.FromErrorToErrorCode(e)
	if err != nil {
		//TODO: Add Domain Error
		//errorCode, err = domainErrors.FromErrorToErrorCode(e)
		return http.StatusInternalServerError, ResponseError{ErrorCode: "UnexpectedError", Message: e.Error()}
	}

	httpStatus, ok := ErrorCodeHttpStatusMap[errorCode]
	if !ok {
		return http.StatusInternalServerError, ResponseError{ErrorCode: "HttpStatusUnmappedError", Message: e.Error()}
	}

	return httpStatus, ResponseError{
		ErrorCode: errorCode.String(),
		Message:   e.Error(),
	}
}
