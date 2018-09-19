package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Job struct {
	url   string
	depth int
}

// History is safe to use concurrently.
type History struct {
	h   map[string]Job
	mux sync.Mutex
}

// Inc increments the history
func (c *History) Add(url string, depth int, queue chan Job) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	_, exist := c.h[url]
	if !exist {
		// fmt.Println(url, depth)
		c.h[url] = Job{url, depth}
		queue <- Job{url, depth}
	}
	c.mux.Unlock()
}

func fetchWeb(url string, depth int, fetcher Fetcher, queue chan Job, successes chan map[string]string, errors chan map[string]string, history *History) {
	if depth <= 0 {
		return
	}
	body, uris, err := fetcher.Fetch(url)
	if err != nil {
		// fmt.Println(err)
		errors <- map[string]string{url: err.Error()}
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)
	successes <- map[string]string{url: body}
	for _, u := range uris {
		// go Crawl(u, depth-1, fetcher)
		history.Add(u, depth-1, queue)
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	queue := make(chan Job)
	successes := make(chan map[string]string)
	errors := make(chan map[string]string)
	history := History{h: make(map[string]Job)}

	quit := time.After(3 * time.Second) // Timeout
	// Fetch URLs in parallel.
	// Don't fetch the same URL twice.
	// This implementation doesn't do either:
	go history.Add("https://golang.org/", 4, queue)
	for {
		select {
		case q := <-queue:
			// fmt.Println(q)
			go fetchWeb(q.url, q.depth, fetcher, queue, successes, errors, &history)
		case s := <-successes:
			for url, body := range s {
				fmt.Println(url, body)
			}
		case e := <-errors:
			for url, err := range e {
				fmt.Println(url, err)
			}
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
