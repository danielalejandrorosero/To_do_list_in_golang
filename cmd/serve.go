package cmd

import (
	"github.com/danielalejandrorosero/to_do_list/app/https"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Start(r *chi.Mux) {

	r.Use(middleware.Logger)
	https.NewSystemController(r)
}
