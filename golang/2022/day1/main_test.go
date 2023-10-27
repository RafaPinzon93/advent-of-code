package main

import "testing"

func TestProcessInput(t *testing.T) {
	test_input := `
    1000
    2000
    3000

    4000

    5000
    6000

    7000
    8000
    9000

    10000
    `

	r := process_input(test_input)
	if r != 24000 {
		t.Errorf("Expected 24000, got %d", r)
	}
}
