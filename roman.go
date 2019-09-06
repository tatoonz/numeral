package numeral

var romanLetterInt = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// RomanToInt converts roman numerals to integer
func RomanToInt(input string) int {
	return romanLetterInt[input]
}
