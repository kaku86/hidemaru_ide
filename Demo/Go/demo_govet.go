package main

import "fmt"

// IDE is the demo structure.
type IDE struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Hidemaru_IDE")
	fmt.Println("Demo_GoVet")

	ide := new(IDE)
	ide.Name = "govet"
}
