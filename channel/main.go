package main

import (
	"fmt"
	"io"
	"net/http"
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
	results := make(chan string)
	for i := 1; i <= 10; i++ {
		url := fmt.Sprintf("http://localhost:9090/page/%d", i)
		go func() {
			results <- fetch(url)
		}()
	}
	bytes := 0
	for range 10 {
		bytes += len(<-results)
	}
	fmt.Println("total bytes:", bytes)
}
