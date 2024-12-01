package main

import (
	"bufio"
	"fmt"
	m "math" // Math library with local alias m.
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	total := 0
	left_ids := make([]int, 0)
	right_ids := make([]int, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		nums := strings.Fields(line)

		if len(nums) == 0 {
			continue
		}

		num, err := strconv.Atoi(nums[0])
		check(err)
		left_ids = append(left_ids, num)

		num, err = strconv.Atoi(nums[1])
		check(err)
		right_ids = append(right_ids, num)
	}

	sort.Ints(left_ids)
	sort.Ints(right_ids)

	for i := 0; i < len(left_ids); i++ {
		total += int(m.Abs(float64(left_ids[i]) - float64(right_ids[i])))
	}

	return total
}

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	total := 0
	left_ids := make([]int, 0)
	right_ids := make([]int, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		nums := strings.Fields(line)

		if len(nums) == 0 {
			continue
		}

		num, err := strconv.Atoi(nums[0])
		check(err)
		left_ids = append(left_ids, num)

		num, err = strconv.Atoi(nums[1])
		check(err)
		right_ids = append(right_ids, num)
	}

	counter := make(map[int]int)

	for _, r_num := range right_ids {
		counter[r_num] += 1
	}

	for _, left_num := range left_ids {
		if val, ok := counter[left_num]; ok {
			total += left_num * val
		}
	}

	return total
}

func main() {
	curr_dir, err := os.Getwd()
	check(err)
	file_dir := curr_dir + "/input.txt"
	dat, err := os.ReadFile(file_dir)
	check(err)

	fmt.Printf("Part 1: %d\n", Part1(string(dat)))
	fmt.Printf("Part 2: %d\n", Part2(string(dat)))
}
