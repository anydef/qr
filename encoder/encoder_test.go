package encoder_test

import (
	"testing"
	"github.com/anydef/qr/encoder"
	"os"
	"encoding/csv"
	"strconv"
)

func Test_DetectMode_Numeric(t *testing.T) {
	mode, _ := encoder.Detect_Mode("1234")
	if mode != encoder.Numeric {
		t.Errorf("Expected nummeric mode: %d", encoder.Numeric)
	}
}

func Test_DetectError_EmptyString(t *testing.T) {
	mode, err := encoder.Detect_Mode("")
	if err == nil {
		t.Errorf("Expected to return error, but got %d", mode)
	}
}

func Test_DetectError_Alphanumeric(t *testing.T) {
	mode, err := encoder.Detect_Mode("12345ABDCEF")
	if err == nil {
		t.Errorf("Expected to return error, but got %d", mode)
	}
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

func parse_mode(record []string, index CsvIndex) encoder.MODE {
	val, err := strconv.Atoi(record[index])
	if err != nil {
		panic(err)
	}
	return encoder.MODE(val)
}

type CharacterCapacities struct {
	Version              int
	ErrorCorrectionLevel rune
	NumericMode          encoder.MODE
	AlphanumericMode     encoder.MODE
	ByteMode             encoder.MODE
	KanjiMode            encoder.MODE
}

func Test_QR_Version(t *testing.T) {
	file, err := os.Open("character_capacities.csv")
	if err != nil {
		t.Fatalf("No character_capacities.csv found. %s", err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		t.Fatalf("error reading all lines: %v", err)
	}
	for i, record := range lines {
		if i == 0 {
			// skip header line
			continue
		}
		version, _ := strconv.Atoi(record[Version])
		errLevel := []rune(record[ErrorCorrectionLevel])[0]
		numericMode := parse_mode(record, NumericMode)
		alphanumericMode := parse_mode(record, AlphanumericMode)
		byteMode := parse_mode(record, ByteMode)
		kanjiMode := parse_mode(record, KanjiMode)
		cc := CharacterCapacities{
			Version:              version,
			ErrorCorrectionLevel: errLevel,
			NumericMode:          numericMode,
			AlphanumericMode:     alphanumericMode,
			ByteMode:             byteMode,
			KanjiMode:            kanjiMode,
		}

		t.Logf("Record %s", cc)
	}

}
