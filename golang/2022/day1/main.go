package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func process_input(input string) (result int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	max_cal := 0
	curr_cal := 0

	for scanner.Scan() {
		cal := strings.TrimSpace(scanner.Text())

		if cal == "" {
			if curr_cal > max_cal {
				max_cal = curr_cal
			}

			curr_cal = 0
		} else {
			cal_int, err := strconv.Atoi(cal)
			check(err)
			curr_cal += cal_int
		}
	}

	return max_cal
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	fmt.Println(process_input(string(dat)))

	// f, err := os.Open("input.txt")
	// check(err)
	// defer f.Close()

	// scanner := bufio.NewScanner(f)

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }
}
