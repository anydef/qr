package encoder

import (
	"strings"
	"errors"
	"strconv"
)

type MODE int

const (
	None    MODE = iota
	Numeric
)

func Detect_Mode(input string) (MODE, error) {
	if len(strings.TrimSpace(input)) == 0 {
		return None, errors.New("Cannot determine encoding mode from input string")
	}

	if _, err := strconv.Atoi(input); err != nil {
		return None, errors.New("Cannot determine encoding mode from input string")
	}
	return Numeric, nil
}
