package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadandParseInput() (reports [][]int) {
	// Open the file
	file, err := os.Open("2024/02/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		tmpReport := make([]int, len(line))
		for _, num := range line {
			parsed, _ := strconv.Atoi(num)
			tmpReport = append(tmpReport, parsed)
		}
		reports = append(reports, tmpReport)
	}

	return
}

func main() {
	reports := ReadandParseInput()
	safe := 0

	fmt.Println("The total number of reports is ", len(reports))

	for _, report := range reports {
		if IsSafeReport(report) {
			safe++
		}
	}

	fmt.Println("The total number of safe reports is ", safe)
}

func IsSafeReport(report []int) bool {
	// TODO(conallob): Debug why this is only reporting 10 safe reports, which is too low
	current, next := 0, 1
	increasing := (report[current] < report[next])

	for current < len(report)-1 && next < len(report) {
		if (report[current] < report[next]) != increasing {
			return false
		}
		if abs(report[current]-report[next]) > 3 {
			return false
		}
		current++
		next++
	}

	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

/*
	if !slices.IsSorted(report) || !IsDecreasing(report) {
		return false
	}
...
func IsDecreasing(report []int) bool {
	return slices.IsSortedFunc(report, func(a, b int) int {
		return cmp.Compare(b, a)
	})
}
*/
