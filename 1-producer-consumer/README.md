# Producer-Consumer Szenario

The producer reads in tweets from a mockstream and a consumer is processing the data to find out whether someone has tweeted about golang or not. The task is to modify the code inside `main.go` so that producer and consumer can run concurrently to increase the throughput of this program.

## Expected results:

Before: 
```
davecheney 	tweets about golang
beertocode 	does not tweet about golang
ironzeb 	tweets about golang
beertocode 	tweets about golang
vampirewalk666 	tweets about golang
Process took 2.499851159s
```

After:
```
davecheney 	tweets about golang
beertocode 	does not tweet about golang
ironzeb 	tweets about golang
beertocode 	tweets about golang
vampirewalk666 	tweets about golang
Process took 1.933252140s
```

# Solution Diff

```
diff --git a/../../go-concurrency-exercises/1-producer-consumer/main.go b/main.go
index 215909c..20b8397 100644
--- a/../../go-concurrency-exercises/1-producer-consumer/main.go
+++ b/main.go
@@ -13,19 +13,20 @@ import (
 	"time"
 )
 
-func producer(stream Stream) (tweets []*Tweet) {
+func producer(stream Stream, tweets chan *Tweet) {
 	for {
 		tweet, err := stream.Next()
 		if err == ErrEOF {
-			return tweets
+			close(tweets)
+			return
 		}
 
-		tweets = append(tweets, tweet)
+		tweets <- tweet
 	}
 }
 
-func consumer(tweets []*Tweet) {
-	for _, t := range tweets {
+func consumer(tweets chan *Tweet) {
+	for t := range tweets {
 		if t.IsTalkingAboutGo() {
 			fmt.Println(t.Username, "\ttweets about golang")
 		} else {
@@ -38,8 +39,10 @@ func main() {
 	start := time.Now()
 	stream := GetMockStream()
 
+	tweets := make(chan *Tweet)
+
 	// Producer
-	tweets := producer(stream)
+	go producer(stream, tweets)
 
 	// Consumer
 	consumer(tweets)
```