package main

import (
	"fmt"
	"testing"
)

func TestProcessInput(t *testing.T) {
	test_input := `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
    `

	fmt.Println(test_input)
	r := process_input(test_input)
	if r != 142 {
		t.Errorf("Expected 142, got %d", r)
	}
}
