package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	ogp "github.com/otiai10/opengraph"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/api/ogp", ogpHandler)
	http.ListenAndServe(":"+port, nil)
}

func ogpHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if len(url) == 0 {
		return
	}
	og, err := ogp.Fetch(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	og = og.ToAbsURL()
	js, err := json.MarshalIndent(og, "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
