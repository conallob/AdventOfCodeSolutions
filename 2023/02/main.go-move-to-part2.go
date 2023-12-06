package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	MAX_CUBES = map[string]int{
		"blue":  14,
		"green": 13,
		"red":   12,
	}
)

func GetGameNumber(line string) string {
	re := regexp.MustCompile(`(\d+):`)
	return re.FindStringSubmatch(line)[1]
}

func IsGamePossible(line string) (int, bool) {
	gameNumber, _ := strconv.Atoi(GetGameNumber(line))

	splitRegexp := regexp.MustCompile(`[\p{P}]`)
	cubes := splitRegexp.Split(line, -1)

	for _, draw := range cubes[1:] {
		drawResult := strings.Split(draw, " ")
		drawCount, err := strconv.Atoi(strings.Trim(drawResult[1], " "))
		if err != nil {
			log.Fatal(err)
		}
		if drawCount > MAX_CUBES[drawResult[2]] {
			return gameNumber, false
		}
	}
	return gameNumber, true

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tally := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameCount, feasible := IsGamePossible(scanner.Text())
		if feasible {
			fmt.Println(scanner.Text())
			tally += gameCount
		}
		
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Final Total: ", tally)
}
