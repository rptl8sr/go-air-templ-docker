package main

import (
	"log"
	"net/http"

	"go-air-templ-docker/views"
)

const (
	defaultAssetsDir = "./public/assets"
)

type Dependencies struct {
	AssetsFS http.FileSystem
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if err := views.Index(views.IndexProps{Name: "INDEX"}).Render(r.Context(), w); err != nil {
		log.Println(err)
	}
}

func handleOther(w http.ResponseWriter, r *http.Request) {
	if err := views.Index(views.IndexProps{Name: "OTHER"}).Render(r.Context(), w); err != nil {
		log.Println(err)
	}
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("OK"))
}

func main() {
	deps := Dependencies{
		AssetsFS: http.Dir(defaultAssetsDir),
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/other", handleOther)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		handler := http.StripPrefix("/assets/", http.FileServer(deps.AssetsFS))
		handler.ServeHTTP(w, r)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
