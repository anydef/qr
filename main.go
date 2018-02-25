package main

import (
	"os"
	"io/ioutil"
	"errors"
)

const output_file_permissions = 0644

func main() {
	if len(os.Args) < 3 {
		panic(errors.New("Missing input parameter"))
	}
	input := os.Args[1]
	path := os.Args[2]

	bytes := []byte(input)
	err := ioutil.WriteFile(path, bytes, output_file_permissions)

	if err != nil {
		panic(err)
	}
}
