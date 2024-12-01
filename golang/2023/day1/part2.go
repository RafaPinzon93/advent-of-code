package part2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_digit_as_string_prefix(s string) (bool, int) {
	nums_string_map := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for key, val := range nums_string_map {
		if strings.HasPrefix(s, key) {
			return true, val
		}
	}

	return false, 0
}

func process_input_2(input string) (result int) {
	/*
		reg := regexp.MustCompile(`\d+`)
		s1 := reg.FindAllString("1abc2asdf6 78d", -1)
		fmt.Println(s1)
		fmt.Println(s1[0])
		fmt.Println(s1[len(s1)-1])
	*/

	scanner := bufio.NewScanner(strings.NewReader(input))
	max_cal := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println(line)
		// for _, char := range line {
		// 	if _, err := strconv.Atoi(string(char)); err == nil {
		// 		fmt.Println(char)

		// 	}
		// }

		val := 0

		for i := 0; i < len(line); i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				val += int(char - '0')
				// } else if digit := get_digit_as_string_prefix(line[i:]) {

			}

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
	fmt.Println(process_input_2(string(dat)))

}
