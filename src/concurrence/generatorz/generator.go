package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - channel only read

func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*)</title>")
			matches := r.FindStringSubmatch(string(html))
			if len(matches) > 0 {
				c <- matches[1]
			} else {
				c <- "" // Ensure we don't try to send invalid data
			}
		}(url)
	}
	return c
}

func main() {
	t1 := title("https://www.google.com", "https://www.amazon.com.br")
	t2 := title("https://www.youtube.com.br", "https://www.github.com.br")

	fmt.Println("First:",  <-t1, "|", t2)
	fmt.Println("Second:",  <-t1, "|", t2)
}