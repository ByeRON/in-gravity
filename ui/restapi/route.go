package restapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"in-gravity/config"
	_ "in-gravity/docs"
	createUser "in-gravity/ui/restapi/user/create"
	"in-gravity/x/ui/restapi/handler"
)

// @title In-Gravity API
// @version 1.0
// @description This is a great application

// @host localhost:3000
// @BasePath /
// @query.collection.format multi
func Serve(conf config.ApplicationConfig) {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/", handler.NewRequestHandler(createUser.NewRequestCreateUserHandleOperator))
	})

	http.ListenAndServe(":3000", r)
}
