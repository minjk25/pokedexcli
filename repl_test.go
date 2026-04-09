package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	test_cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello   World    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Harry   Potter    ",
			expected: []string{"harry", "potter"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, test_case := range test_cases {
		actual := cleanInput(test_case.input)
		if len(actual) != len(test_case.expected) {
			t.Errorf("cleanInput(%q) returned %v, expected %v", test_case.input, actual, test_case.expected)
			continue
		}

		for index, word := range actual {
			if word != test_case.expected[index] {
				t.Errorf("cleanInput(%q) returned %v, expected %v", test_case.input, actual, test_case.expected)
			}

		}
	}
}
