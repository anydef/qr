package parsing_test

import (
	"testing"
	"github.com/anydef/qr/parsing"
	"fmt"
)

func Test_Version_Numeric_Low(t *testing.T) {
	input := "123456"
	version := parsing.Get_Version(input, parsing.Low)

	fmt.Println(version)

	if version.Capacity != 41 {
		t.Fatalf("")
	}

	if version.Ordinal != 1 {
		t.Fatalf("")
	}
}

func Test_Version_Numeric_Medium(t *testing.T) {
	input := "123456"
	version := parsing.Get_Version(input, parsing.Medium)

	fmt.Println(version)

	if version.Capacity != 34 {
		t.Fatalf("")
	}

	if version.Ordinal != 1 {
		t.Fatalf("")
	}
}
