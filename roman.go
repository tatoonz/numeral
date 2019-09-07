package numeral

var romanLetterInt = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// RomanToInt converts roman numerals to integer
func RomanToInt(input string) int {
	result := 0
	prevX := 0
	for _, letter := range input {
		x := romanLetterInt[letter]

		if prevX < x {
			result = (result - prevX) + (x - prevX)
		} else {
			result += romanLetterInt[letter]
		}
		prevX = x
	}

	return result
}
