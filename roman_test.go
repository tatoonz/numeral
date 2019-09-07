package numeral_test

import (
	"fmt"
	"testing"

	"github.com/tatoonz/numeral"
)

func TestRomanToInt(t *testing.T) {
	t.Run("OneLetter", func(t *testing.T) {
		tests := []struct {
			input    string
			expected int
		}{
			{input: "I", expected: 1},
			{input: "V", expected: 5},
			{input: "X", expected: 10},
			{input: "L", expected: 50},
			{input: "C", expected: 100},
			{input: "D", expected: 500},
			{input: "M", expected: 1000},
		}

		for _, test := range tests {
			actual := numeral.RomanToInt(test.input)
			if actual != test.expected {
				t.Errorf("expected %s to be %d, but got %d", test.input, test.expected, actual)
			}
		}
	})

	t.Run("MultipleSameLetters", func(t *testing.T) {
		tests := []struct {
			input    string
			expected int
		}{
			{input: "II", expected: 2},
			{input: "III", expected: 3},
			{input: "XX", expected: 20},
			{input: "XXX", expected: 30},
			{input: "CC", expected: 200},
			{input: "CCC", expected: 300},
			{input: "MM", expected: 2000},
			{input: "MMM", expected: 3000},
		}

		for _, test := range tests {
			actual := numeral.RomanToInt(test.input)
			if actual != test.expected {
				t.Errorf("expected %s to be %d, but got %d", test.input, test.expected, actual)
			}
		}
	})

	t.Run("Subtract", func(t *testing.T) {
		tests := []struct {
			input    string
			expected int
		}{
			{input: "IV", expected: 4},
			{input: "IX", expected: 9},
			{input: "XL", expected: 40},
			{input: "XC", expected: 90},
			{input: "CD", expected: 400},
			{input: "CM", expected: 900},
		}

		for _, test := range tests {
			actual := numeral.RomanToInt(test.input)
			if actual != test.expected {
				t.Errorf("expected %s to be %d, but got %d", test.input, test.expected, actual)
			}
		}
	})

	t.Run("ComplexNumerals", func(t *testing.T) {
		tests := []struct {
			input    string
			expected int
		}{
			{input: "MMMCCXC", expected: 3290},
			{input: "CXXIII", expected: 123},
			{input: "DCCLXXVII", expected: 777},
			{input: "CMXLIX", expected: 949},
			{input: "MM", expected: 2000},
			{input: "LXXXVII", expected: 87},
			{input: "XLIII", expected: 43},
			{input: "XXII", expected: 22},
			{input: "DCCVII", expected: 707},
			{input: "LXIX", expected: 69},
			{input: "XCIX", expected: 99},
			{input: "MDLXXXII", expected: 1582},
			{input: "MMXIX", expected: 2019},
		}

		for _, test := range tests {
			actual := numeral.RomanToInt(test.input)
			if actual != test.expected {
				t.Errorf("expected %s to be %d, but got %d", test.input, test.expected, actual)
			}
		}
	})
}

func ExampleRomanToInt() {
	fmt.Println(numeral.RomanToInt("I"))
	fmt.Println(numeral.RomanToInt("IV"))
	fmt.Println(numeral.RomanToInt("XI"))
	fmt.Println(numeral.RomanToInt("MMXIX"))
	// output:
	// 1
	// 4
	// 11
	// 2019
}
