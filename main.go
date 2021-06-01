package main

import (
  "fmt"

  "test.com/go-project/morestrings"
  "github.com/google/go-cmp/cmp"
)

func main() {
  fmt.Println(morestrings.ReverseRunes("Hello World!"))
  fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
