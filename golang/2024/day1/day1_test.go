package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example input",
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			expected: 11,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)
			if got != tt.expected {
				t.Errorf("Part1() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example input",
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			expected: 31, // Replace with actual expected value
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.input)
			if got != tt.expected {
				t.Errorf("Part2() = %v, want %v", got, tt.expected)
			}
		})
	}
}
