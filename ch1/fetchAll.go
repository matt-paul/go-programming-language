package main

// input

// ➜ ./fetchAll https://golang.org http://gopl.io https://godoc.org

// output....

// 1.34s    8377 https://golang.org
// 1.69s    6805 https://godoc.org
// 4.43s    4154 http://gopl.io
// 4.43s e;apsed

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// create a channel of strings using make
	ch := make(chan string)
	// for each command line argument, start a goroutine that calls
	// fetch asynchronously to getch url using http.Get
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs e;apsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	// io.Copy reads the body of the response, and discards it by writing to the ioutil.Discard output stream.
	// Copy returns the byte count
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // dont leak
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	// As each result arrives, fetch sends a summary line on the channel ch, which the second range loop in main recieves and prints
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// input

// ➜ ./fetchAll https://golang.org http://gopl.io https://godoc.org

// output....

// 1.34s    8377 https://golang.org
// 1.69s    6805 https://godoc.org
// 4.43s    4154 http://gopl.io
// 4.43s e;apsed
