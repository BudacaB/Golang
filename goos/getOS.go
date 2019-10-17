package main

import (
	"fmt"
	"runtime"
)

func execute() {

	const GOARCH string = runtime.GOARCH
	const GOOS string = runtime.GOOS

	fmt.Println(GOARCH)
	fmt.Println(GOOS)

}

func main() {
	execute()
}
