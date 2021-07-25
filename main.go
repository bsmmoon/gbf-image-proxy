package main

import (
	"log"
	"net/http"

	"gbf-image-proxy/requests"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/get/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[len("/get/"):]
	w.Write(requests.Get(url))
}
