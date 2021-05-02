package create

import (
	"net/http"

	"in-gravity/applications"
	"in-gravity/x/ui/restapi/handler"
)

type RequestCreateUserHandleOperator struct {
	presenter applications.CreateUserPresenterInterface
	writer    http.ResponseWriter
	request   *http.Request
}

func NewRequestCreateUserHandleOperator(writer http.ResponseWriter, request *http.Request) handler.RequestHandleOperatorInterface {
	return &RequestCreateUserHandleOperator{
		writer:  writer,
		request: request,
	}
}

func (o *RequestCreateUserHandleOperator) SetupPresenter() error {
	o.presenter = NewRestCreateUserPresenter(o.writer)
	return nil
}

func (o *RequestCreateUserHandleOperator) ExecuteService() {
}
