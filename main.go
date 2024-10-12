package main

import (
	"log"
	"net/http"

	"go-air-templ-docker/views"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if err := views.Index(views.IndexProps{Name: "INDEX"}).Render(r.Context(), w); err != nil {
		log.Println(err)
	}
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/health", handleHealth)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
