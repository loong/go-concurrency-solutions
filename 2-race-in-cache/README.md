# Race condition in caching szenario

Given is some code to cache key-value pairs from a mock database into
the main memory (to reduce access time). The code is buggy and
contains a race conditition. Change the code to make this thread safe.

*Note*: that golang's map are not entirely thread safe. Multiple readers
are fine, but multiple writers are not.

# Test your solution

Use the following command to test for race conditions:
```
go run -race *.go
```

Correct solution:
No output = solution correct:
```
$ go run -race *.go
$
```

Incorrect solution:
```
==================
WARNING: DATA RACE
Write by goroutine 7:
...
==================
Found 3 data race(s)
```

# Solution Diff

Note that we **only need the lock when we attempt to write** to the cache, since reads from map[string] are threadsafe.

```
diff --git a/../../go-concurrency-exercises/2-race-in-cache/main.go b/main.go
index 33557a3..fec18c4 100644
--- a/../../go-concurrency-exercises/2-race-in-cache/main.go
+++ b/main.go
@@ -8,7 +8,10 @@
 
 package main
 
-import "container/list"
+import (
+	"container/list"
+	"sync"
+)
 
 // CacheSize determines how big the cache can grow
 const CacheSize = 100
@@ -24,6 +27,7 @@ type KeyStoreCache struct {
 	cache map[string]string
 	pages list.List
 	load  func(string) string
+	lock  sync.Mutex
 }
 
 // New creates a new KeyStoreCache
@@ -40,6 +44,7 @@ func (k *KeyStoreCache) Get(key string) string {
 
 	// Miss - load from database and save it in cache
 	if !ok {
+		k.lock.Lock()
 		val = k.load(key)
 		k.pages.PushFront(key)
 
@@ -48,6 +53,7 @@ func (k *KeyStoreCache) Get(key string) string {
 			delete(k.cache, k.pages.Back().Value.(string))
 			k.pages.Remove(k.pages.Back())
 		}
+		k.lock.Unlock()
 	}
 
 	return val
```