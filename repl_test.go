package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello WORlD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "      heLlo  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  world ",
			expected: []string{"world"},
		},
		{
			input:    " hello world . ",
			expected: []string{"hello", "world", "."},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The length of the actual slice (%d) does not match the length of the expected slice (%d)", len(actual), len(c.expected))
			t.Errorf("Actual slice: %v", actual)
			t.Errorf("Expected slice: %v", c.expected)
			return
		}

		for i, actualWord := range actual {
			expectedWord := c.expected[i]
			if expectedWord != actualWord {
				t.Errorf("The expected and actual are not the same: %s != %s", expectedWord, actualWord)
				return
			}
		}
	}
}
