package main

import (
	"fmt"
	"runtime"
)

// CurrentVersion is the current application's version literal
const CurrentVersion = "0.1.0-dev"

// APIVersion is the current API version
const APIVersion = "/igcfs/" + CurrentVersion + "/"

// printVersion prints the current version
func printVersion() {
	fmt.Printf("igcfs (go) version: %s\n", CurrentVersion)
	fmt.Printf("System version: %s\n", runtime.GOARCH+"/"+runtime.GOOS)
	fmt.Printf("Go version: %s\n\n", runtime.Version())
}
