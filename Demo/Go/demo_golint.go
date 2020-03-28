package main

import "fmt"

var (
	version string
)

// Version returns the version
func Version() string {
	return version
}

func main() {
	fmt.Println("Hidemaru_IDE")
	fmt.Println("Demo_Golint")

}
