package parsing

type InputType int

var constLookup = map[InputType]string{
	Numeric:      `Numeric`,
	Alphanumeric: `Alphanumeric`,
	Byte:         `Byte`,
	Kanji:        `Kanji`,
}

const (
	_            InputType = iota
	Numeric
	Alphanumeric
	Byte
	Kanji
)

func (i InputType) String() string {
	return constLookup[i]
}

func DetermineInputType(input string) InputType {
	r := Numeric
	for _, c := range input {
		if char_alphanumeric(c) && r <= Alphanumeric {
			r = Alphanumeric
		}
		if char_byte(c) && r <= Byte {
			r = Byte
		}
		if c > 255 && r <= Kanji {
			r = Kanji
		}
	}
	return r
}

func char_byte(c rune) bool {
	return c > 96 && c <= 255
}
func char_alphanumeric(c rune) bool {
	return c > 35 && c < 48 || c > 57 && c <= 96
}
