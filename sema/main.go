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
	sem := make(chan struct{}, 3) // at most three requests at a time
	for i := 1; i <= 10; i++ {
		url := fmt.Sprintf("http://localhost:9090/page/%d", i)
		wg.Go(func() {
			sem <- struct{}{}        // take a slot, wait if all three are taken
			defer func() { <-sem }() // give it back
			fmt.Println(fetch(url))
		})
	}
	wg.Wait()
}
