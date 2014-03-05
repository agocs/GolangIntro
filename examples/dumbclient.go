package main

import (
	"fmt"
	"net/http"
	"time"
)

type httpResp struct {
	who       string
	resp      string
	timeTaken time.Duration
}

func main() {

	respChan := make(chan httpResp)

	urlList := []string{"http://example.com", "http://tumblr.com", "http://yahoo.com", "http://penny-arcade.com", "http://google.com", "http://apple.com", "http://python.org", "http://mitsubishi.com", "http://toyota.com", "http://ford.com", "http://chevrolet.com", "http://audi.com", "http://subaru.com", "http://bmw.com", "http://honda.com", "http://motorola.com", "http://cingular.com", "http://att.com", "http://ibm.com", "http://compaq.com", "http://dell.com", "http://cdw.com"}

	urlsTried := 0
	startTime := time.Now()
	for _, url := range urlList {
		urlsTried++
		go bother(url, respChan)
	}

	for i := 0; i < urlsTried; i++ {
		resp := <-respChan
		time.Sleep(1000000000)
		fmt.Printf("tried %s, got %s in %s time\n", resp.who, resp.resp, resp.timeTaken)
	}

	fmt.Println(time.Since(startTime))
}

func bother(who string, respChan chan httpResp) {
	startTime := time.Now()

	resp, err := http.Get(who)
	newResponse := httpResp{"", "", time.Since(startTime)}

	if err != nil {
		newResponse = httpResp{who, err.Error(), time.Since(startTime)}
	} else {
		newResponse = httpResp{who, resp.Status, time.Since(startTime)}
	}
	respChan <- newResponse
	newResponse.who = "terry"
	fmt.Println("Done bothering")
}
