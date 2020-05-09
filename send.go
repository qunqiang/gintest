package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	Threads = 10
	Jobs    = 1000
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(Threads)
	for i := 0; i < Threads; i++ {
		go func(jobId int) {
			for j := 0; j < Jobs; j++ {
				msg := fmt.Sprintf("%d-%d", jobId, j)
				sendRequest(msg)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

}

func sendRequest(msg string) {
	client := &http.Client{}
	body := bufio.NewReader(bytes.NewBufferString(msg))
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/ping", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("traceid", msg)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	respContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respContent))
}
