package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: "hello", expected: []string{"hello"}},
		{input: "hello world", expected: []string{"hello", "world"}},
		{input: "hello\nworld", expected: []string{"hello", "world"}},
		{input: "hello\tworld", expected: []string{"hello", "world"}},
		{input: "hello\rworld", expected: []string{"hello", "world"}},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("expected: %v, got: %v", cs.expected, actual)
			continue
		}
		for i := range actual {
			if actual[i] != cs.expected[i] {
				t.Errorf("expected: %v, got: %v", cs.expected, actual)
				break
			}
		}
	}

}
