package main

import (
	"log"
	"sync"
	"time"
)

type VisitedURLs struct {
	urls map[string]bool
	mux  sync.Mutex
}

var rateLimiter = time.Tick(500 * time.Millisecond)

func main() {
	log.Println("Starting Concurrent Web Crawler")
	startURL := "https://www.w3schools.com/"
	urls := make(chan string)
	visited := &VisitedURLs{urls: make(map[string]bool)}
	var wg sync.WaitGroup

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, urls, visited)
	}

	urls <- startURL

	// go func() {
	// 	time.Sleep(20 * time.Second)
	// 	close(urls)
	// }()

	wg.Wait()
}

func (v *VisitedURLs) IsVisited(url string) bool {
	v.mux.Lock()
	defer v.mux.Unlock()
	return v.urls[url]
}

func (v *VisitedURLs) Add(url string) {
	v.mux.Lock()
	defer v.mux.Unlock()
	v.urls[url] = true
}

func worker(id int, wg *sync.WaitGroup, urls chan string, visited *VisitedURLs) {
	defer wg.Done()
	<-rateLimiter
	for url := range urls {
		if visited.IsVisited(url) {
			continue
		}

		visited.Add(url)
		log.Printf("Worker %d fetching URL: %s\n", id, url)

		if !canFetch(url) {
			continue
		}

		content, err := fetch(url)
		if err != nil {
			log.Println("Error fetching page:", err)
			continue
		}

		links, err := parseLinks(content)
		if err != nil {
			log.Println("Error parsing page:", err)
			continue
		}
		for _, link := range links {
			urls <- link
		}
	}
}
