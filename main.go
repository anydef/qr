package main

import (
	"io/ioutil"
	"flag"
	"fmt"
)

const output_file_permissions = 0644

var correction_level string
var input string
var output string
var mask_pattern int

func parse_varargs() {
	flag.StringVar(&correction_level, "correction-level", "low", "error correction level")
	flag.IntVar(&mask_pattern, "mask-pattern", 0, "mask pattern. TBD")
	flag.StringVar(&input, "input", "", "Input string")
	flag.StringVar(&output, "output", "", "Output file path")

	flag.Parse()
}

func main() {
	parse_varargs()
	bytes := []byte(input)
	err := ioutil.WriteFile(output, bytes, output_file_permissions)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Input string %s\n", input)
	fmt.Printf("Correction level %s\n", correction_level)
	fmt.Printf("Mask pattern %d\n", mask_pattern)
	fmt.Printf("Output path %s\n", output)

}
