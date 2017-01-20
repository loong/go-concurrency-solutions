# Limit your crawler

Given is a crawler (modified from the Go tour), which would pull pages excessively. The task is to modify the `main.go` file, so that

 - the crawler cannot crawl more than 2 pages per second
 - without losing it's concurrency

# Test your solution

Use `go test` to verify if your solution is correct.

Correct solution:
```
PASS
ok      github.com/mindworker/go-concurrency-exercises/0-limit-crawler  13.009s
```

Incorrect solution:
```
--- FAIL: TestMain (7.80s)
        main_test.go:18: There exists a two crawls who were executed less than 1 sec apart.
	        main_test.go:19: Solution is incorrect.
		FAIL
		exit status 1
		FAIL    github.com/mindworker/go-concurrency-exercises/0-limit-crawler  7.808s
```

# Solution Diff

```
diff --git a/../../go-concurrency-exercises/0-limit-crawler/main.go b/main.go
index f2cdfa7..aa84677 100644
--- a/../../go-concurrency-exercises/0-limit-crawler/main.go
+++ b/main.go
@@ -11,8 +11,11 @@ package main
 import (
 	"fmt"
 	"sync"
+	"time"
 )
 
+var limiter <-chan time.Time
+
 // Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a
 // real crawler. It crawls until the maximum depth has reached.
 func Crawl(url string, depth int, wg *sync.WaitGroup) {
@@ -22,6 +25,8 @@ func Crawl(url string, depth int, wg *sync.WaitGroup) {
 		return
 	}
 
+	<-limiter
+
 	body, urls, err := fetcher.Fetch(url)
 	if err != nil {
 		fmt.Println(err)
@@ -40,6 +45,8 @@ func Crawl(url string, depth int, wg *sync.WaitGroup) {
 func main() {
 	var wg sync.WaitGroup
 
+	limiter = time.Tick(1 * time.Second)
+
 	wg.Add(1)
 	Crawl("http://golang.org/", 4, &wg)
 	wg.Wait()
```