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

	if len(nums) >= 1 {
	
		score := 1

		for i := 0 ; i < len(nums) ; i++ {
			score *= 2
		}
		
		return score
	} else {
		return 0
	}
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
		score = ScoreWinningScratchcard(slices.Clip(winningNums))

		fmt.Println(winningNums, " - ", score)

		tally += score

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Final Tally: ", tally)

}
