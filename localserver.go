package main

import (
  "fmt"
  "net/http" // https://golang.org/pkg/net/http/
)

// リクエストの処理(クロージャー的にもかける）
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}

func main() {
  http.HandleFunc("/", handler) // '/'に対して第二引数で処理(handler)を追加
  http.ListenAndServe(":8080", nil) // 8080ポートで立ち上げ
}
