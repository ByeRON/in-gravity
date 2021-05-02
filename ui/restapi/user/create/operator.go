package create

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"in-gravity/applications"
	"in-gravity/x/ui/restapi/handler"
)

type RequestCreateUserHandleOperator struct {
	presenter applications.CreateUserPresenterInterface
	service   applications.CreateUserApplicationServiceInterface
	input     applications.CreateUserInput
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

func (o *RequestCreateUserHandleOperator) SetupService() error {
	// appconf

	//TODO: use wire_gen
	o.service = applications.NewCreateUserApplicationService(o.presenter)
	return nil
}

func (o *RequestCreateUserHandleOperator) ParseInput() error {
	//OK: Get request
	//OK: Unmarshal request and transfer RequestCreateUser
	//OK: transfer RequestCreateUser to input
	//NG: API definition -> swaggo

	body := o.request.Body
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, body)
	if err != nil {
		o.presenter.OutputError(err)
		return err
	}

	var request RequestCreateUser
	err = json.Unmarshal(buf.Bytes(), &request)
	if err != nil {
		o.presenter.OutputError(err)
		return err
	}

	input := applications.CreateUserInput{
		Name: string(request.Name),
	}

	o.input = input
	return nil
}

func (o *RequestCreateUserHandleOperator) ExecuteService() {
	o.service.Handle()
}
