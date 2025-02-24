package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "my test",
			expected: []string{"my", "test"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "hell0 wORld",
			expected: []string{"hell0", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Slice lenght mismatch: Actual: %v | Expected: %v ", len(actual), len(c.expected))
		}
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Input: %s | Expected: %s | Actual: %s", c.input, expectedWord, word)
			}
		}
	}
}
