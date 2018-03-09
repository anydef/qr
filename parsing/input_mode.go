package parsing

type InputMode int

var inputModeToString = map[InputMode]string{
	Numeric:      `Numeric`,
	Alphanumeric: `Alphanumeric`,
	Byte:         `Byte`,
	Kanji:        `Kanji`,
}

const (
	_            InputMode = iota
	Numeric
	Alphanumeric
	Byte
	Kanji
)

func (i InputMode) String() string {
	return inputModeToString[i]
}

func DetermineInputType(input string) InputMode {
	r := Numeric
	for _, c := range input {
		if char_alphanumeric(c) && r <= Alphanumeric {
			r = Alphanumeric
		}
		if char_byte(c) && r <= Byte {
			r = Byte
		}
		if char_kanji(c) && r <= Kanji {
			return Kanji
		}
	}
	return r
}
func char_kanji(c rune) bool {
	return c > 255
}

func char_byte(c rune) bool {
	return c > 96 && c <= 255
}
func char_alphanumeric(c rune) bool {
	return c > 35 && c < 48 || c > 57 && c <= 96
}
