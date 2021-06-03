package requests

import (
  "io/ioutil"
  "log"
  "net/http"
  "fmt"
)

func Get(url string) []byte {
  url = fmt.Sprintf("http://%s", url)

  resp, err := http.Get(url)
  if err != nil {
    log.Fatalln(err)
  }

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatalln(err)
  }

  return body
}

