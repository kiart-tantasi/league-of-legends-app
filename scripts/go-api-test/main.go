package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Duration(5000 * time.Millisecond),
	}
	checkHealthEndpoint(&client)
	checkMatchesEndpoint(&client)
	fmt.Println("==================\nAll tests passed!\n==================")
}

func checkHealthEndpoint(client *http.Client) {
	isSuccess := false
	for i := 0; i < 5; i++ {
		isSuccessChannel := make(chan bool)
		go func() {
			res, err := client.Get("http://localhost:8080/api/health")
			if err == nil && res.StatusCode == 200 {
				isSuccessChannel <- true
			} else {
				isSuccessChannel <- false
			}
		}()
		isSuccess = <-isSuccessChannel
		if isSuccess {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if !isSuccess {
		panic("health failed")
	}
}

func checkMatchesEndpoint(client *http.Client) {
	res, err := client.Get("http://localhost:8080/api/v1/matches?gameName=GAMENAME&tagLine=TAGLINE")
	if err != nil || res.StatusCode != 200 {
		panic("matches failed")
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	fmt.Println(string(bytes))
}
