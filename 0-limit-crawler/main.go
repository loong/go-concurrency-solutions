//////////////////////////////////////////////////////////////////////
//
// Your task is to change the code to limit the crawler so that it is
// not possible fo the crawler to crawl more than one page per second.
//
// @hint: you can archive that by adding 3 lines
//

package main

import (
	"fmt"
	"sync"
	"time"
)

var limiter <-chan time.Time

// Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a
// real crawler. It crawls until the maximum depth has reached.
func Crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	<-limiter

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	for _, u := range urls {
		go Crawl(u, depth-1, wg)
	}
	return
}

func main() {
	var wg sync.WaitGroup

	limiter = time.Tick(1 * time.Second)

	wg.Add(1)
	Crawl("http://golang.org/", 4, &wg)
	wg.Wait()
}
