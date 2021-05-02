package create

import (
	"in-gravity/applications"
	"net/http"
)

type RestCreateUserPresenter struct {
	w http.ResponseWriter
}

func NewRestCreateUserPresenter(w http.ResponseWriter) applications.CreateUserPresenterInterface {
	return RestCreateUserPresenter{w: w}
}

func (p RestCreateUserPresenter) Output() {
	return
}

func (p RestCreateUserPresenter) OutputError(e error) {
	return
}
