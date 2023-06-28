package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Start Server")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World!"))
		if err != nil {
			log.Println("Error Warning : ", err)
		}
	})
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println("Fail Server : ", err)
	}
}
