package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/get/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[len("/get/"):]
	url = fmt.Sprintf("http://%s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	io.Copy(w, resp.Body)
}
