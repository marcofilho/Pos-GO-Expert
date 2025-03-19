package main

import (
	"fmt"

	"github.com/marcofilho/fcutils-secret/pkg/events"
)

func main() {
	eventDispatcher := events.NewEventDispatcher()
	fmt.Println(eventDispatcher)
}
