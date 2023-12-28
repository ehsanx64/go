package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	defaultUrl = "https://api.namefake.com/"
)

type ReqResult struct {
	Body     string
	Request  *http.Request
	Response *http.Response
}

func httpGet(url string) ReqResult {
	verbose := false
	out := ReqResult{}
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	out.Body = string(body)
	out.Request = req
	out.Response = res

	if verbose {
		if res.StatusCode != http.StatusOK {
			fmt.Println("Non-200 status code detected")
		}

		fmt.Printf("res: %v\n", string(body))
	}

	return out
}

// This map will contain the urls
var urls = make(map[string]string)

// Add a url by appending input string to the base (default) url
func addUrl(u string) {
	urls[u] = defaultUrl + u
}

func main() {
	// Add some urls
	addUrl("1")
	addUrl("2")
	addUrl("3")

	// Iterate over urls map and fetch all urls
	for k, v := range urls {
		var input map[string]interface{}

		fmt.Println("<=", k)
		req := httpGet(v)
		fmt.Println("!!", "Received", len(req.Body), "bytes")

		if err := json.Unmarshal([]byte(req.Body), &input); err != nil {
			log.Println("Error unmarshalling data for url:", v)
			log.Println(err)
		}

		fmt.Println("=>", "Name:", input["name"])
	}
}
