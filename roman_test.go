package numeral_test

import (
	"testing"

	"github.com/tatoonz/numeral"
)

func TestConvertRomanToInt(t *testing.T) {
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
}
