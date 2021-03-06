package parsing

import (
	"fmt"
	"encoding/csv"
	"io"
	"os"
	"errors"
	"strconv"
)

type CorrectionLevel rune

var correction_level_ToString = map[CorrectionLevel]string{
	Low:     `Low`,
	Medium:  `Medium`,
	Quality: `Quality`,
	High:    `High`,
}

var string_to_level = map[string]CorrectionLevel{
	`L`: Low,
	`M`: Medium,
	`Q`: Quality,
	`H`: High,
}

func CorrectionLevelFromString(s string) CorrectionLevel {
	return string_to_level[s]
}

func (c CorrectionLevel) String() string {
	return correction_level_ToString[c]
}

const (
	_       CorrectionLevel = iota
	Low
	Medium
	Quality
	High
)

type Version struct {
	Capacity int
	Ordinal  int
}

func Get_Version(input string, level CorrectionLevel) Version {
	fmt.Printf("Input [%s], correction level [%s]", input, level)

	capacity, ordinal := find_version(DetermineInputType(input), level, len(input))
	return Version{Capacity: capacity, Ordinal: ordinal}
}

func find_version(mode InputMode, level CorrectionLevel, input_len int) (int, int) {
	const (
		version          int = iota
		error_correction
		numeric
		alphanumeric
		byte_mode
		kanji
	)

	file, err := os.Open("character_capacities.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)
	skip_header(r)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if level != CorrectionLevelFromString(record[error_correction]) {
			continue
		}

		version := to_int(record[version])
		switch mode {
		case Numeric:
			if c := to_int(record[numeric]); input_len <= c {
				return c, version
			}
			continue
		case Alphanumeric:
			if c := to_int(record[alphanumeric]); input_len <= c {
				return c, version
			}
			continue
		case Byte:
			if c := to_int(record[byte_mode]); input_len <= c {
				return c, version
			}
			continue
		case Kanji:
			if c := to_int(record[kanji]); input_len <= c {
				return c, version
			}
			continue
		}
	}
	panic(errors.New(fmt.Sprintf("Couldn't find encoding record for input mode [%s] and level [%s]", mode, level)))
}

func skip_header(r *csv.Reader) {
	_, err := r.Read()
	if err != nil {
		panic(err)
	}
}

func to_int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Invalid value [%s]. Cannot parse to int", s)))
	}
	return i
}
