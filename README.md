# Go Concurrency: goroutines, WaitGroups and channels

Code for the video [Go Concurrency in 5 Minutes: Goroutines, WaitGroups and Channels](https://youtu.be/SOLcD9oHVfY) and the post [Go concurrency, measured](https://huseyinbabal.dev/posts/go-concurrency-goroutines-waitgroup-channels/).

Every client talks to one local server that sleeps one second per request, so ten requests in a row take ten seconds and anything faster comes from doing them at the same time.

```bash
go run ./server &        # localhost:9090, 1 s per request
```

| Step | Run | What it shows | Measured (Go 1.27.1, M4 Max) |
|---|---|---|---|
| plain loop | `time go run ./sequential` | ten requests one after another | 10.32 s |
| `go` keyword only | `time go run ./gokeyword` | `main` returns before any goroutine prints | 0.33 s, no output |
| `sync.WaitGroup` + `wg.Go` | `time go run ./waitgroup` | all ten in parallel, random order | 1.27 s |
| shared counter | `go run -race ./race` | data race on `bytes += n` | `Found 2 data race(s)` |
| channel | `go run -race ./channel` | results through a channel, no race | 101 bytes |
| fan-out / fan-in | `time go run ./fanout` | three workers, jobs + results channels | 4.32 s |
| buffered-channel limit | `time go run ./sema` | at most three requests at a time | 4.37 s |
| goroutine cost | `go run ./cost` | 100,000 blocked goroutines | ~2 KB stack each |

`wg.Go` needs Go 1.25 or newer.
