package main

import (
  "fmt"

  "github.com/google/go-cmp/cmp"

  "test.com/go-project/morestrings"
  "test.com/go-project/requests"
)

func main() {
  fmt.Println(morestrings.ReverseRunes("Hello World!"))
  fmt.Println(cmp.Diff("Hello World", "Hello Go"))
  requests.Get("")
}
