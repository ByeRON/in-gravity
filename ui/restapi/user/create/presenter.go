package create

import (
	"encoding/json"
	"in-gravity/applications"
	"in-gravity/ui/restapi/defs"
	"net/http"
)

type RestCreateUserPresenter struct {
	w http.ResponseWriter
}

func NewRestCreateUserPresenter(w http.ResponseWriter) applications.CreateUserPresenterInterface {
	return RestCreateUserPresenter{w: w}
}

<<<<<<< HEAD
func (p RestCreateUserPresenter) Output() {
=======
func (p RestCreateUserPresenter) Output(o applications.CreateUserOutput) {
	p.w.WriteHeader(http.StatusAccepted)
	data := ResponseCreateUser{
		UserId: o.UserId,
	}
	body, _ := json.Marshal(data)
	p.w.Write(body)
>>>>>>> Add
}

func (p RestCreateUserPresenter) OutputError(e error) {
	var statusCode int
	statusCode, structResponse := defs.NewResponseError(e)
	res, _ := json.Marshal(structResponse)
	p.w.WriteHeader(statusCode)
	p.w.Write(res)
}
