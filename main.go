package main

import (
	"encoding/json"
	"fmt"
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
		return
	}
	og = og.ToAbsURL()
	b, err := json.MarshalIndent(og, "", "\t")
	fmt.Fprintf(w, "%+v", string(b))
}
