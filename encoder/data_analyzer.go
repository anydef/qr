package encoder

import (
	"strings"
	"errors"
	"strconv"
	"encoding/csv"
	"os"
	"fmt"
)

type MODE int

const (
	None    MODE = iota
	Numeric
)

type CorrectionLevel int

const (
	_       CorrectionLevel = iota
	Low
	Medium
	Quality
	High
)

func Detect_Mode(input string) (MODE, error) {
	if len(strings.TrimSpace(input)) == 0 {
		return None, errors.New("Cannot determine encoding parse_mode from input string")
	}

	if _, err := strconv.Atoi(input); err != nil {
		return None, errors.New("Cannot determine encoding parse_mode from input string")
	}
	return Numeric, nil
}

type CsvIndex int

const (
	Version              CsvIndex = iota
	ErrorCorrectionLevel
	NumericMode
	AlphanumericMode
	ByteMode
	KanjiMode
)

func parse_mode(record []string, index CsvIndex) int {
	val, err := strconv.Atoi(record[index])
	if err != nil {
		panic(err)
	}
	return val
}

type VersionCapacities struct {
	Version              int
	ErrorCorrectionLevel rune
	NumericMode          int
	AlphanumericMode     int
	ByteMode             int
	KanjiMode            int
}

func (cc VersionCapacities) Size() int {
	return 17 + (cc.Version * 4)
}

func parseVersionCapacity(record []string) VersionCapacities {
	version, _ := strconv.Atoi(record[Version])
	errLevel := []rune(record[ErrorCorrectionLevel])[0]

	return VersionCapacities{
		Version:              version,
		ErrorCorrectionLevel: errLevel,
		NumericMode:          parse_mode(record, NumericMode),
		AlphanumericMode:     parse_mode(record, AlphanumericMode),
		ByteMode:             parse_mode(record, ByteMode),
		KanjiMode:            parse_mode(record, KanjiMode),
	}
}

func (c *Context) DetermineVersion(input string, level CorrectionLevel) (int, error) {

	if len(input) == 0 {
		return 1, nil
	}
	max_version := c.Capacities[len(c.Capacities)-1]

	if level == Low && len(input) > max_version.NumericMode ||
		level == Medium && len(input) > 3391 ||
		level == Quality && len(input) > 2420 ||
		level == High && len(input) > 1852 {
		return 0, errors.New(fmt.Sprintf("Input len( %d ) too long.", len(input)))
	}

	return 40, nil
}

type Context struct {
	Capacities []VersionCapacities
}

var _context *Context = nil

func GetContext() *Context {
	if _context == nil {
		_context = &Context{Capacities: loadCapacitiesTable()}
	}
	return _context
}

func loadCapacitiesTable() []VersionCapacities {
	file, err := os.Open("character_capacities.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	capacities := make([]VersionCapacities, len(lines)-1)

	for i, record := range lines {
		if i == 0 {
			continue
		}
		capacities[i-1] = parseVersionCapacity(record)
	}
	return capacities
}
