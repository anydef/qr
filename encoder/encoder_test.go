package encoder_test

import (
	"testing"
	"github.com/anydef/qr/encoder"
	//"sort"
	//"fmt"
)

func Test_DetectMode_Numeric(t *testing.T) {
	mode, _ := encoder.Determine_InputType("1234")
	if mode != encoder.Numeric {
		t.Errorf("Expected nummeric mode: %d", encoder.Numeric)
	}
}

func Test_DetectError_EmptyString(t *testing.T) {
	mode, err := encoder.Determine_InputType("")
	if err == nil {
		t.Errorf("Expected to return error, but got %d", mode)
	}
}

func _Test_DetectError_Alphanumeric(t *testing.T) {

	mode, err := encoder.Determine_InputType("12345ABDCEF")
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

		str := make_string(limit_size, '1')

		version, err := context.DetermineVersion(str, level)

		if err != nil {
			t.Fatalf("%d should fit into numeric capacity for correction level %d", limit_size, level)
		}

		if version != 40 {
			t.Fatalf("input len(%d) for correction level %s should return version 40", limit_size, level)
		}

	}

}
func make_string(len int, char rune) string {
	chars := make([]rune, len)
	for i := range chars {
		chars[i] = char
	}
	return string(chars)
}

func Test_Find_SmallestVersion(t *testing.T) {
	context := encoder.GetContext()
	if v := context.Smallest_Numeric_Version_L(make_string(10, '1')); v != 1 {
		t.Fatalf("Wrong version returned. [ %d ]", v)
	}

	if v := context.Smallest_Numeric_Version_L(make_string(42, '1')); v != 2 {
		t.Fatalf("Wrong version returned. [ %d ]", v)
	}

	//type pair struct {
	//	version  int
	//	capacity int
	//}
	//
	//a := []pair{{1, 100}, {3, 300}, {6, 600}, {10, 1000}}
	//x := 101
	//
	//i := sort.Search(len(a), func(i int) bool {
	//	return a[i].capacity >= x
	//})
	//
	//if i < len(a) && a[i].capacity >= x {
	//	fmt.Printf("found%d at index %d in %v\n", x, i, a)
	//	fmt.Printf("%v", a[i])
	//} else {
	//	fmt.Printf("%d not found in %v\n", x, a)
	//}
}
