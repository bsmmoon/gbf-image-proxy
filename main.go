package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

	log.Print(fmt.Sprintf("Received %s", url))

	if !checkSafety(url) {
		log.Print("Failed safety check.")
		w.WriteHeader(400)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}

	io.Copy(w, resp.Body)
}

func checkSafety(url string) bool {
	flag := true
	flag = flag && strings.HasPrefix(url, "http://game-a.granbluefantasy.jp")
	return flag
}
