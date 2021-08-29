package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/get/", handler)
	log.Print(fmt.Sprintf("Listening to %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
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
