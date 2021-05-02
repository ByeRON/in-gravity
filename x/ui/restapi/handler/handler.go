package handler

import "net/http"

type NewRequestHandleOperator func(http.ResponseWriter, *http.Request) RequestHandleOperatorInterface

func NewRequestHandler(f NewRequestHandleOperator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		operator := f(w, r)
		manager := RequestHandleManager{Operator: operator}
		manager.Handle()
	}
}
