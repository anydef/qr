package encoder_test

import (
	"testing"
	"github.com/anydef/qr/encoder"
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

type pair struct {
	l, r interface{}
}

func Test_QR_Version(t *testing.T) {
	context := encoder.GetContext()

	for _, version_capacity := range context.Capacities {
		size := 17 + (version_capacity.Version * 4)
		if version_capacity.Size() != size {
			t.Fatalf("Size for version %d should be %d", version_capacity.Version, size)
		}
	}
}

func Test_DetermineVersion_Empty_All_CorrectionLevels(t *testing.T) {
	context := encoder.GetContext()
	levels := []encoder.CorrectionLevel{encoder.Low, encoder.Medium, encoder.Quality, encoder.High}
	for _, level := range levels {
		i, err := context.DetermineVersion("", level)
		if err != nil {
			t.Fatalf("Should be verison 1 for empty string, got: %s", err)
		}
		if i != 1 {
			t.Fatalf("Version for empty string should be 1")
		}
	}
}



func Test_DetermineVersion_Numeric_OutOfBounds_CorrectionLevel_Low(t *testing.T) {
	context := encoder.GetContext()

	size := 1853

	chars := make([]rune, size)

	for i := range chars {
		chars[i] = '1'
	}
	level := encoder.Low
	_, err := context.DetermineVersion(string(chars), level)
	t.Logf(err.Error())
	if err == nil {
		t.Fatalf("%d should not fit into numeric capacity for %s", size, level)
	}

}
