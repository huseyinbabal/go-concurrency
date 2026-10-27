package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func main() {
	var wg sync.WaitGroup
	bytes := 0
	for i := 1; i <= 10; i++ {
		url := fmt.Sprintf("http://localhost:9090/page/%d", i)
		wg.Go(func() {
			bytes += len(fetch(url))
		})
	}
	wg.Wait()
	fmt.Println("total bytes:", bytes)
}
