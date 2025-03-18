package main

import "fmt"

func main() {
	events := []string{"birthday", "party", "wedding", "graduation"}
	fmt.Println("First version of events", events)

	events = append(events[:0], events[1:]...)

	fmt.Println("Second version of events", events)
}
