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
		if i != 0 {
			fmt.Println("health endpoint retry round", i)
		}
		res, err := client.Get("http://localhost:8080/api/health")
		if err == nil && res.StatusCode == 200 {
			isSuccess = true
			break
		}
		time.Sleep(2 * time.Second)
	}
	if !isSuccess {
		panic("health endpoint failed")
	}
}

func checkMatchesEndpoint(client *http.Client) {
	res, err := client.Get("http://localhost:8080/api/v1/matches?gameName=GAMENAME&tagLine=TAGLINE")
	if err != nil || res.StatusCode != 200 || res.Body == nil {
		panic("matches endpoint failed")
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic("matches endpoint failed")
	}
	fmt.Println(string(bytes))
}
