package parsing_test

import (
	"testing"
	"github.com/anydef/qr/parsing"
)

func Test_DetermineInput_empty(t *testing.T) {
	if parsing.DetermineInputType("") != parsing.Numeric {
		t.Fatalf("Should be numeric for empty string")
	}
}

func Test_DetermineInput_Numbers(t *testing.T) {
	tests := []string{
		"1", "0123456789",
	}
	expected := parsing.Numeric
	for _, test := range tests {
		if r := parsing.DetermineInputType(test); r != expected {
			t.Fatalf("[%s] Should be [%s], but was [%s] \n", test, expected, r)
		}
	}
}

func Test_DetermineInput_Alphanumeric(t *testing.T) {
	tests := []string{
		"0123456789$", "A", "Z", "$%*+-.,/ ", "B",
	}
	expected := parsing.Alphanumeric
	for _, test := range tests {
		if r := parsing.DetermineInputType(test); r != expected {
			t.Fatalf("[%s] Should be [%s], but was [%s] \n", test, expected, r)
		}
	}

}

func Test_DetermineInput_Byte(t *testing.T) {
	tests := []string{
		"absd",
	}
	expected := parsing.Byte
	for _, test := range tests {
		if r := parsing.DetermineInputType(test); r != expected {
			t.Fatalf("[%s] Should be [%s], but was [%s] \n", test, expected, r)
		}
	}
}

func Test_DetermineInput_Kanji(t *testing.T) {
	tests := []string{
		"ワ",
	}
	expected := parsing.Kanji
	for _, test := range tests {
		if r := parsing.DetermineInputType(test); r != expected {
			t.Fatalf("[%s] Should be [%s], but was [%s] \n", test, expected, r)
		}
	}

}
