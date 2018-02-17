package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		input := os.Args[1]
		fmt.Printf("QR generated: %s\n", input)
	} else {
		fmt.Println("Missing input paramter")
	}
}
