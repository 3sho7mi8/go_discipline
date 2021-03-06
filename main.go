package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	uri := "http://localhost:8080"

	maxConnection := make(chan bool, 200)
	wg := &sync.WaitGroup{}

	count := 0
	start := time.Now()
	for maxRequest := 0; maxRequest < 10000; maxRequest++ {
		wg.Add(1)
		maxConnection <- true
		go func() {
			defer wg.Done()

			resp, err := http.Get(uri)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			count++
			<-maxConnection
		}()
	}
	wg.Wait()
	end := time.Now()
	log.Printf("%d 回のリクエストに成功しました!\n", count)
	log.Printf("%f 秒処理に時間がかかりました!\n", (end.Sub(start)).Seconds())
}
