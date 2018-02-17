package encoder

import (
	"testing"
	"strings"
	"errors"
	"strconv"
)

type MODE int

const (
	None    MODE = iota
	Numeric
)

func detect_mode(input string) (MODE, error) {
	if len(strings.TrimSpace(input)) == 0 {
		return None, errors.New("Cannot determine encoding mode from input string")
	}

	if _, err:= strconv.Atoi(input); err != nil {
		return None, errors.New("Cannot determine encoding mode from input string")
	}
	return Numeric, nil
}

func TestDetected_ModeNumeric(t *testing.T) {
	mode, _ := detect_mode("1234")
	if mode != Numeric {
		t.Errorf("Expected nummeric mode: %d", Numeric)
	}
}

func TestDetected_ErrorEmptyString(t *testing.T) {
	mode, err := detect_mode("")
	if err == nil {
		t.Errorf("Expected to return error, but got %d", mode)
	}
}

func TestDetected_ErrorAlphanumeric(t *testing.T) {
	mode, err := detect_mode("12345ABDCEF")
	if err == nil {
		t.Errorf("Expected to return error, but got %d", mode)
	}
}
