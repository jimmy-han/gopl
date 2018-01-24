package main

import (
	"fmt"
	"log"
	"net/http"
)

func request(hostname string) (response string) {
	resp, err := http.Get(hostname)
	if err != nil {
		log.Fatal(err)
	}
	return hostname + ":" + resp.Status
}

func mirroredQuery(responses chan string) {
	go func() { responses <- request("http://www.baidu.com") }()
	go func() { responses <- request("http://usa.baidu.com") }()
	go func() { responses <- request("http://www.baidu.jp") }()
}

func main() {
	var responses = make(chan string, 3)
	mirroredQuery(responses)
	for response := range responses {
		fmt.Println(response)
	}
	close(responses)
}
