package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ReadAndParseInput() (matches [][2]int) {
	// Open the file
	file, err := os.Open("2024/03/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	pattern := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)

	for scanner.Scan() {
		// TODO(conallob): Finish this function to parse the input file
		stringMatches := pattern.FindAllString(scanner.Text(), -1)
		lhs, err := strconv.Atoi(stringMatches[0])
		if err != nil {
			continue
		}
		rhs, err := strconv.Atoi(stringMatches[1])
		if err != nil {
			continue
		}
		// TODO(conallob): Replace []int{} objects w/ a struct
		matches = append(matches, [2]int{lhs, rhs})
	}
	fmt.Println(matches)
	return
}

func main() {
	instructions := ReadAndParseInput()
	tally := 0

	for _, instruction := range instructions {
		tally += instruction[0] * instruction[1]
	}

	fmt.Println("The total is:", tally)
}
