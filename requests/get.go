package requests

import (
  "io/ioutil"
  "log"
  "net/http"
)

func Get(url string) {
  example := "http://game-a1.granbluefantasy.jp/assets_en/img_mid/sp/assets/npc/gacha/3030310000_01.png"
  resp, err := http.Get(example)
  if err != nil {
    log.Fatalln(err)
  }

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatalln(err)
  }

  sb := string(body)
  log.Printf(sb)
}

