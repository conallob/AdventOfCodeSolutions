package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/juliangruber/go-intersect"
)


func ScoreWinningScratchcard(nums []interface{}) int {

	// TODO(conallob): Fix scoring issue in this function
	return 2^(len(nums))
}


func main() {
	file, err := os.Open("example-part1.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
  re := regexp.MustCompile(`[[:punct:]]`)

	tally := 0
	score := 0

	for scanner.Scan() {
  	parsed := re.Split(scanner.Text(), -1)
		score = 0

		winners := strings.Split(parsed[1], " ")
		scratchcard := strings.Split(parsed[2], " ")

		winningNums := slices.DeleteFunc(
        intersect.Simple(winners, scratchcard),
        func(thing interface{}) bool {
            return thing == nil
        },
    )
		score = ScoreWinningScratchcard(winningNums)

		fmt.Println(winningNums, " - ", score)

		tally += score

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Final Tally: ", tally)

}
