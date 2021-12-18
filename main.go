package main

import (
	"fmt"

	"github.com/0x0844/go_generate_example/types"
)


func main() {
	var event types.Event = 0
	fmt.Println(event)	// Add
	event = 1
	fmt.Println(event)	// Sub
}