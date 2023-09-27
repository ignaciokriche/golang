// Exercise: Web Crawler
//
// In this exercise you'll use Go's concurrency features to parallelize a web crawler.
//
// Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
//
// Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not
// safe for concurrent use!

package main

import (
	"fmt"
	"sync"
)

func ExerciseWebCrawlerDemo() {

	fmt.Println("Exercise Web Crawler")

	Crawl("https://golang.org/", 4, fetcher)

	fmt.Println("----------------------------------------------------")
	fmt.Println()

}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// use a private function to keep track of visited urls:
	crawl(url, depth, fetcher, makeUrlMutexVisitor())
}

func crawl(url string, depth int, fetcher Fetcher, visitor *urlMutexVisitor) {

	if depth <= 0 {
		return
	}

	if visitor.visited(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\tfound: %s %q\n", url, body)

	wg := sync.WaitGroup{}
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			crawl(url, depth-1, fetcher, visitor)
		}(u)
	}
	wg.Wait()
}

type urlMutexVisitor struct {
	visitUrlMutex sync.Mutex
	knownUrls     map[string]bool
}

func makeUrlMutexVisitor() *urlMutexVisitor {
	return &urlMutexVisitor{sync.Mutex{}, make(map[string]bool)}
}

func (v *urlMutexVisitor) visited(url string) bool {

	v.visitUrlMutex.Lock()
	defer v.visitUrlMutex.Unlock()

	if _, ok := v.knownUrls[url]; ok {
		return true
	}
	v.knownUrls[url] = true
	return false
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
	return "", nil, fmt.Errorf("\tnot found: %s", url)
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
