package main

import (
	"fmt"
)

func main() {
	printVersion()
	options := processCmdLn()
	h, c, err := initiateHost(options)
	defer c()

	if err != nil {
		panic(err)
	}

	fmt.Printf("My hosts ID is %s  --   %s\n", h.ID().Pretty(), h.ID())

	// do nothing, just wait for the kill
	fmt.Scanln()
}
