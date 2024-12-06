package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadAndParseInput() (reports [][]int) {
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
		for i, num := range line {
			tmpReport[i], _ = strconv.Atoi(num)
		}
		reports = append(reports, tmpReport)
	}

	return
}

func main() {
	reports := ReadAndParseInput()
	safe, partial := 0, 0

	fmt.Println("The total number of reports is ", len(reports))

	for _, report := range reports {
		if IsSafeReport(report) {
			safe++
		} else if IsPartiallySafeReport(report) {
			partial++
		}
	}

	fmt.Println("The total number of safe reports is ", safe)

	fmt.Println("The total number of partially safe reports is ", safe+partial)
}

func IsSafeReport(report []int) bool {
	current, next := 0, 1
	if report[current] == report[next] {
		return false
	}
	direction := report[current] < report[next]
	unsafeCount := 0

	for (current < len(report)-1) && (next < len(report)) {
		diff := abs(report[current] - report[next])
		if ((report[current] < report[next]) != direction) || diff < 1 || diff > 3 {
			unsafeCount++
		}
		current++
		next++
	}
	return unsafeCount == 0
}

func IsPartiallySafeReport(report []int) bool {
	/* TODO(conallob): Debug why this function isn't
	   working as intended
	*/
	subtotal := 0

	fmt.Println("Report: ", report)
	i := 0
	permutation := make([]int, len(report)-1)
	for i < len(report) {
		if i == 0 {
			permutation = report[1:]
		} else if i == len(report)-1 {
			permutation = report[:len(report)-1]
		} else {
			permutation = append(report[:i], report[i+1:]...)
		}
		fmt.Println(permutation)
		if IsSafeReport(permutation) {
			subtotal++
		}
		i++
	}
	return (len(report) - subtotal) == 1
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
