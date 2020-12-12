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
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Need at least 2 argument <Method> and <URL>")
		os.Exit(1)
	}
	curlAll(os.Args[1], os.Args[2:])
}

func curlAll(method string, urls []string) {
	waitChannel := make(chan (string))
	for _, url := range urls {
		go curl(method, url, waitChannel)
	}
	for _, url := range urls {
		fmt.Println(url, <-waitChannel)
	}
}

func curl(method, url string, waitChannel chan<- string) {
	now := time.Now()
	curlClient := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		waitChannel <- fmt.Sprint("could not create a request")
		return
	}

	resp, err := curlClient.Do(req)
	if err != nil {
		waitChannel <- fmt.Sprint("could not execute a request")
		return
	}

	dataSize, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		waitChannel <- fmt.Sprint("Could not count response data size.")
		return
	}

	secs := time.Since(now)
	waitChannel <- fmt.Sprintf("received %vKbs response body after %.4f secs", dataSize/1024, secs.Seconds())
}
