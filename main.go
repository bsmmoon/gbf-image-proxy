package main

import (
  "net/http"
  "log"

  "test.com/go-project/requests"
)

func main() {
  log.SetFlags(log.LstdFlags | log.Lshortfile)

  http.HandleFunc("/get/", handler)
  log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  url := r.URL.Path[len("/get/"):]
  w.Write(requests.Get(url))
}

