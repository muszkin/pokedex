package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input              string
		expected           []string
		expectedReturnSize int
	}{
		{
			input:              "  hello  world  ",
			expected:           []string{"hello", "world"},
			expectedReturnSize: 2,
		}, {
			input:              "TY to Te ",
			expected:           []string{"ty", "to", "te"},
			expectedReturnSize: 3,
		}, {
			input:              "PiKaChU iS SUPER",
			expected:           []string{"pikachu", "is", "super"},
			expectedReturnSize: 3,
		}, {
			input:              "8799yv97chsh9348904r20u24",
			expected:           []string{"8799yv97chsh9348904r20u24"},
			expectedReturnSize: 1,
		}, {
			input:              "Y^(*&(*& U(",
			expected:           []string{"y^(*&(*&", "u("},
			expectedReturnSize: 2,
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != c.expectedReturnSize {
			t.Errorf("return slice different from expected input: %v, output: %v", c.expectedReturnSize, len(actual))
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("return value different from expected input: %v, output %v", expectedWord, word)
				return
			}
		}
	}
}
