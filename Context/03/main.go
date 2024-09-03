package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, error := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if error != nil {
		return
	}

	res, error := http.DefaultClient.Do(req)
	if error != nil {
		return
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
