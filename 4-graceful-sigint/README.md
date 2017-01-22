# Graceful SIGINT killing

Given is a mock process which runs indefinitely and blocks the program. Right now the only way to stop the program is to send a SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
want to try to gracefully stop the process first.

Change the program to do the following:
   1. On SIGINT try to gracefully stop the process using
          `proc.Stop()`
   2. If SIGINT is called again, just kill the program (last resort)

# Solution Diff

```
diff --git a/../../go-concurrency-exercises/4-graceful-sigint/main.go b/main.go
index 38c7d3d..6e1dfba 100644
--- a/../../go-concurrency-exercises/4-graceful-sigint/main.go
+++ b/main.go
@@ -13,10 +13,26 @@
 
 package main
 
+import (
+	"os"
+	"os/signal"
+)
+
 func main() {
 	// Create a process
 	proc := MockProcess{}
 
+	// Listen to signal
+	c := make(chan os.Signal, 1)
+	signal.Notify(c, os.Interrupt)
+
+	go func() {
+		<-c
+		go proc.Stop()
+		<-c
+		os.Exit(0)
+	}()
+
 	// Run the process (blocking)
 	proc.Run()
 }
```