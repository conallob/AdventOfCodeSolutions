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

func CubesPower(cubes map[string]int) int {
	out := (cubes["red"] * cubes["green"] * cubes["blue"])
	fmt.Println(out)
	return out
}

func GetGameNumber(line string) string {
	re := regexp.MustCompile(`(\d+):`)
	return re.FindStringSubmatch(line)[1]
}

func IsGamePossible(line string, max_cubes map[string]int) (int, bool) {
	gameNumber, _ := strconv.Atoi(GetGameNumber(line))

	splitRegexp := regexp.MustCompile(`[\p{P}]`)
	cubes := splitRegexp.Split(line, -1)

	for _, draw := range cubes[1:] {
		drawResult := strings.Split(draw, " ")
		drawCount, err := strconv.Atoi(strings.Trim(drawResult[1], " "))
		if err != nil {
			log.Fatal(err)
		}
		if drawCount > max_cubes[drawResult[2]] {
			max_cubes[drawResult[2]] = drawCount
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
		max_cubes := map[string]int{
			"blue":  1,
			"green": 1,
			"red":   1,
		}

		gameDetails := scanner.Text()
		_, feasible := IsGamePossible(gameDetails, max_cubes)
		if feasible {
			fmt.Println(scanner.Text(), " - ", max_cubes)
			tally += CubesPower(max_cubes)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Final Total: ", tally)
}
