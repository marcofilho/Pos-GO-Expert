package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Duration(time.Microsecond),
	}

	resp, err := client.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Body:", string(body))
}
