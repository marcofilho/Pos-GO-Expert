package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}

	json := bytes.NewBuffer([]byte(`{"Name":"Marco"}`))

	resp, err := client.Post("https://google.com", "application/json", json)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
