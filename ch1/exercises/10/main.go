// Request a website twice to see if there is any difference in response time
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	ch := make(chan string)
	url := os.Args[1]

	go fetch(url, ch)
	fmt.Println(<-ch)

	go fetch(url, ch)
	fmt.Println(<-ch)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("[%s] %.2fs  %7d  %s", resp.Status, secs, nbytes, url)
}
