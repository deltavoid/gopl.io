// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
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

	fmt.Println("main: 1")
	start := time.Now()
	
	ch := make(chan string)
	
    fmt.Println("main: 2")
	for _, url := range os.Args[1:] {


	    fmt.Println("main: 4")
		go fetch(url, ch) // start a goroutine
	
		fmt.Println("main: 5")
	}
	
	fmt.Println("main: 6")
	for range os.Args[1:] {
	
		fmt.Println("main: 7")
		fmt.Println(<-ch) // receive from channel ch
	
		fmt.Println("main: 8")
	}
	
	fmt.Println("main: 9")
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

    fmt.Println("main: 10")
}

func fetch(url string, ch chan<- string) {
	
	fmt.Println("fetch: 1")
	start := time.Now()

	fmt.Println("fetch: 2")
	resp, err := http.Get(url)
	
	fmt.Println("fetch: 3")
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	fmt.Println("fetch: 4")
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	
	fmt.Println("fetch: 5")
	secs := time.Since(start).Seconds()
	
	fmt.Println("fetch: 6")
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
	
	fmt.Println("fetch: 7")
}

//!-
