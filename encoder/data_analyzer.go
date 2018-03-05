package encoder

import (
	"strings"
	"errors"
	"strconv"
	"encoding/csv"
	"os"
	"fmt"
)

type InputType int

const (
	None         InputType = iota
	Numeric
	Alphanumeric
	Byte
	Kanji
)

type CorrectionLevel int

const (
	_       CorrectionLevel = iota
	Low
	Medium
	Quality
	High
)

var CorrectionLevelMap map[rune]CorrectionLevel = map[rune]CorrectionLevel{
	'L': Low,
	'M': Medium,
	'Q': Quality,
	'H': High,
}

func Determine_InputType(input string) (InputType, error) {
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
	NumericField
	AlphanumericField
	ByteField
	KanjiField
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
		NumericMode:          parse_mode(record, NumericField),
		AlphanumericMode:     parse_mode(record, AlphanumericField),
		ByteMode:             parse_mode(record, ByteField),
		KanjiMode:            parse_mode(record, KanjiField),
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

type VersionCapacity struct {
	version  int
	capacity int
}

type Context struct {
	Capacities      []VersionCapacities
	capacity_lookup map[InputType]map[CorrectionLevel][]VersionCapacity
}

func (c *Context) Smallest_Numeric_Version_L(input string) int {
	input_type, err := Determine_InputType(input)
	if err != nil {
		panic(err)
	}
	return c.capacity_lookup[input_type][Low]
}

var _context *Context = nil

func GetContext() *Context {
	if _context == nil {
		table := loadCapacitiesTable()
		type pair struct {
			version  int
			capacity int
		}
		////var types map[string]int
		////types = make(map[string]int)
		//low := []pair{}
		//medium := []pair{}
		//quality := []pair{}
		//high := []pair{}
		capacity_lookup := make(map[InputType]map[CorrectionLevel][]VersionCapacity)
		for _, capacity := range table {
			versionCapacities := capacity_lookup[Numeric][CorrectionLevelMap[capacity.ErrorCorrectionLevel]]
			versionCapacities = append(versionCapacities, VersionCapacity{capacity.Version, capacity.NumericMode})
		}

		_context = &Context{Capacities: table, capacity_lookup: capacity_lookup}
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
