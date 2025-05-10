package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello","world"},
		},
		{
			input: "this is atest",
			expected: []string{"this","is","atest"},
		},
		{
			input: "",
			expected: []string{},
		},
	}


	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Length mismatch! Expected: %d | Actual: %d", len(c.expected), len(actual))
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word mismatch at %d! Expected : %s | Actual : %s", i, expectedWord, word)
				return
			}
		}
	}
}