package main

import (
	"fmt"
	"net/http"
	"time"
)

type httpResponse struct {
	url         string
	status      string
	timeElapsed time.Duration
}

func main() {

	respChan := make(chan httpResponse)

	startTime := time.Now()
	urlsTried := 0
	urlList := []string{"http://example.com", "http://tumblr.com", "http://yahoo.com", "http://penny-arcade.com", "http://google.com", "http://apple.com", "http://python.org", "http://mitsubishi.com", "http://toyota.com", "http://ford.com", "http://chevrolet.com", "http://audi.com", "http://subaru.com", "http://bmw.com", "http://honda.com", "http://motorola.com", "http://cingular.com", "http://att.com", "http://ibm.com", "http://compaq.com", "http://dell.com", "http://cdw.com"}

	for _, url := range urlList {
		go bother(url, respChan)
		urlsTried++
	}

	for i := 0; i < urlsTried; i++ {
		resp := <-respChan
		fmt.Printf("%s returned %s in %s\n", resp.url, resp.status, resp.timeElapsed)
	}

	fmt.Println(time.Since(startTime))
}

func bother(who string, respChan chan httpResponse) {

	startTime := time.Now()

	resp, err := http.Get(who)

	if err != nil {
		respChan <- httpResponse{who, err.Error(), time.Since(startTime)}
	}
	respChan <- httpResponse{who, resp.Status, time.Since(startTime)}

}
