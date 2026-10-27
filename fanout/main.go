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

func worker(id int, jobs <-chan string, results chan<- string) {
	for url := range jobs {
		results <- fmt.Sprintf("worker %d: %s", id, fetch(url))
	}
}

func main() {
	jobs := make(chan string)
	results := make(chan string)

	// fan-out: three workers read from the same jobs channel
	var wg sync.WaitGroup
	for id := 1; id <= 3; id++ {
		wg.Go(func() { worker(id, jobs, results) })
	}

	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- fmt.Sprintf("http://localhost:9090/page/%d", i)
		}
		close(jobs)
	}()

	// fan-in: close results once every worker is done
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}
