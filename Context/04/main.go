package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	showToken(ctx)
}

func showToken(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
