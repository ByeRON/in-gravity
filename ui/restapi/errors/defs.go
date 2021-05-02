package errors

import "golang.org/x/xerrors"

type UiErrorCode string

const (
	ErrorSomething UiErrorCode = "Ui.Something"
)

var something = xerrors.New(ErrorSomething.String())

var uiErrorCodeErrorMap = map[UiErrorCode]error{
	ErrorSomething: something,
}

func (e UiErrorCode) String() string {
	return string(e)
}

func FromErrorToErrorCode(e error) (UiErrorCode, error) {
	if xerrors.Is(e, something) {
		return ErrorSomething, nil
	}
	var notMatched UiErrorCode
	return notMatched, xerrors.New("Matched error is nothing")
}
