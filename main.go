package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		index, err := ioutil.ReadFile("./index.html")
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		w.Write(index)
	})

	r.Post("/{username}", func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		fmt.Println(username)
		// Get request to username.voiceit.io
	})

	log.Fatal(http.ListenAndServe(":3000", r))
}
