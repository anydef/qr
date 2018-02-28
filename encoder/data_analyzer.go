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
	Low     CorrectionLevel = iota
	Medium
	Quality
	High
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

type CsvIndex int

const (
	Version              CsvIndex = iota
	ErrorCorrectionLevel
	NumericMode
	AlphanumericMode
	ByteMode
	KanjiMode
)

func mode(record []string, index CsvIndex) MODE {
	val, err := strconv.Atoi(record[index])
	if err != nil {
		panic(err)
	}
	return MODE(val)
}

type VersionCapacities struct {
	Version              int
	ErrorCorrectionLevel rune
	NumericMode          MODE
	AlphanumericMode     MODE
	ByteMode             MODE
	KanjiMode            MODE
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
		NumericMode:          mode(record, NumericMode),
		AlphanumericMode:     mode(record, AlphanumericMode),
		ByteMode:             mode(record, ByteMode),
		KanjiMode:            mode(record, KanjiMode),
	}
}

func (c *Context) DetermineVersion(input string, level CorrectionLevel) (int, error) {
	if len(input) > 1852 {
		return 0, errors.New(fmt.Sprintf("Input len( %d ) too long.", len(input)))
	}
	return 1, nil
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
