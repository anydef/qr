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

func _Test_DetectError_Alphanumeric(t *testing.T) {
	mode, err := encoder.Detect_Mode("12345ABDCEF")
	if err == nil {
		t.Errorf("Expected to return error, but got %d", mode)
	}
	t.Fatalf("Not implemeted")
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

func Test_DetermineVersion_Numeric_Above_MaxValue(t *testing.T) {
	size_limits := []pair{
		{encoder.Low, 4296 + 1},
		{encoder.Medium, 3391 + 1},
		{encoder.Quality, 2420 + 1},
		{encoder.High, 1852 + 1},
	}
	context := encoder.GetContext()

	for _, limit := range size_limits {
		level := limit.l.(encoder.CorrectionLevel)
		limit_size := limit.r.(int)

		chars := make([]rune, limit_size)

		for i := range chars {
			chars[i] = '1'
		}
		_, err := context.DetermineVersion(string(chars), level)

		if err == nil {
			t.Fatalf("%d should not fit into numeric capacity for correction level %d", limit_size, level)
		}

		t.Logf(err.Error())

	}

}

func Test_DetermineVersion_Numeric_MaxValue(t *testing.T) {
	size_limits := []pair{
		{encoder.Low, 4296},
		{encoder.Medium, 3391},
		{encoder.Quality, 2420},
		{encoder.High, 1852},
	}
	context := encoder.GetContext()

	for _, limit := range size_limits {
		level := limit.l.(encoder.CorrectionLevel)
		limit_size := limit.r.(int)

		chars := make([]rune, limit_size)

		for i := range chars {
			chars[i] = '1'
		}
		version, err := context.DetermineVersion(string(chars), level)

		if err != nil {
			t.Fatalf("%d should fit into numeric capacity for correction level %d", limit_size, level)
		}

		if version != 40 {
			t.Fatalf("input len(%d) for correction level %s should return version 40", limit_size, level)
		}

	}

}
