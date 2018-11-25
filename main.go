package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
)

func main() {
  // アクセスするためのURI
  uri := "http://localhost:8080"

  // GETリクエストでアクセス
  // err でエラー時にエラーの内容が入り、panicで処理を中断してエラーの中身を出力
  // deferで関数が終了するとクローズするおまじない
  resp, err := http.Get(uri)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  // レスポンスを取り出す処理
  byteArray, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(byteArray))
}
