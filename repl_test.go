package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	testCases := []TestCase{
		{
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "  Go   is   AWESOME  ",
			expected: []string{"go", "is", "awesome"},
		},
		{
			input:    "singleword",
			expected: []string{"singleword"},
		},
		{
			input:    "MiXeD   CaSe   WoRdS",
			expected: []string{"mixed", "case", "words"},
		},
	}

	for _, testCase := range testCases {
		actual := cleanInput(testCase.input)

		if !reflect.DeepEqual(testCase.expected, actual) {
			t.Fatalf("expected: %v, got: %v", testCase.expected, actual)
		}
	}
}
