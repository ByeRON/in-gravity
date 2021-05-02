package restapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"in-gravity/config"
	createUser "in-gravity/ui/restapi/user/create"
	"in-gravity/x/ui/restapi/handler"
)

func Serve(conf config.ApplicationConfig) {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/", handler.NewRequestHandler(createUser.NewRequestCreateUserHandleOperator))
	})

	http.ListenAndServe(":3000", r)
}
