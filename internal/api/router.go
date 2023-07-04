package api

import (
	"github.com/go-chi/chi"
)

func main(controller *APIController) {
	router := chi.NewRouter()

	router.Route("/user", func(r chi.Router) {
		r.Post("/", controller.Create)
		r.Get("/", controller.GetAll)
		r.Get("/userId}", controller.GetSingle)
		r.Put("/{userId}", controller.Update)
		r.Delete("/{userId}", controller.Delete)
	})
}
