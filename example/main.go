package main

import (
	"fmt"
	"net/http"

	almost "github.com/anstk/almost-router"
)

func main() {

	r := almost.Router()

	r.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Index page"))
	})

	r.Route("GET", "/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User page"))
	})

	r.Route("GET", "/play", play)

	r.Start("localhost:8000")

}

func play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "played")
}
