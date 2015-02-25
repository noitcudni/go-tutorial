package main

import (
	"fmt"
)

type UrlInfo struct {
	url   string
	depth int
}

var (
	crawled_url = map[string]int{}
	add_chan    = make(chan string)
	wait        = make(chan int)
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) <-chan string {
	c := make(chan string)

	go func() {
		if depth <= 0 {
			close(c)
			return
		}
		_, urls, err := fetcher.Fetch(url)
		if err != nil {
			//fmt.Println(err, depth)
			close(c)
			return
		}

		//crawled_url[url] = depth
		//fmt.Printf("found: %s \n", url, depth)
		add_chan <- url
		<-wait

		for _, u := range urls {
			<-Crawl(u, depth-1, fetcher)
		}
		c <- url
	}()

	return c
}

func link_bookkeeper() {
	for {
		u := <-add_chan
		_, ok := crawled_url[u]
		if ok == false {
			crawled_url[u] = 0
		}
		wait <- 0
	}
}

func main() {
	go link_bookkeeper()
	c := Crawl("http://golang.org/", 4, fetcher)
	<-c

	fmt.Println(crawled_url)
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
