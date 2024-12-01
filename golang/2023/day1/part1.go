package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func process_input(input string) (result int) {
	/*
		reg := regexp.MustCompile(`\d+`)
		s1 := reg.FindAllString("1abc2asdf6 78d", -1)
		fmt.Println(s1)
		fmt.Println(s1[0])
		fmt.Println(s1[len(s1)-1])
	*/

	scanner := bufio.NewScanner(strings.NewReader(input))
	reg := regexp.MustCompile(`\d`)
	max_cal := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println(line)
		nums := reg.FindAllString(line, -1)
		fmt.Println(nums)
		// for _, char := range line {
		// 	if _, err := strconv.Atoi(string(char)); err == nil {
		// 		fmt.Println(char)

		// 	}
		// }

		val := 0

		if len(nums) == 0 {
			continue
		} else if len(nums) == 1 {
			num, err := strconv.Atoi(nums[0] + nums[0])
			check(err)
			val = num
		} else {
			num, err := strconv.Atoi(string(nums[0] + nums[len(nums)-1]))
			check(err)
			val = num
		}

		max_cal += val
	}

	return max_cal
}

func main() {
	curr_dir, err := os.Getwd()
	check(err)
	file_dir := curr_dir + "/golang/2023/day1/input.txt"
	dat, err := os.ReadFile(file_dir)
	check(err)
	fmt.Println(process_input(string(dat)))

}
