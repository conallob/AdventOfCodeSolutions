package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/juliangruber/go-intersect"
)


func ScoreWinningScratchcard(nums []interface{}) int {

	var sanitisedNums []interface{}

	// Filter out nil values
	for _, a := range nums {
		if ((a == "") || (a == nil) || (a == " ")) {
			continue
		} else {
			sanitisedNums = append(sanitisedNums, a)
		}
	}

	if len(sanitisedNums) == 1 {
		return 1
	} else {
		return int(math.Pow(2, float64(len(sanitisedNums) -1)))
	}
}


func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
  	re := regexp.MustCompile(`[[:punct:]]`)

	tally := 0

	for scanner.Scan() {
  	parsed := re.Split(scanner.Text(), -1)
		score := 0

		winners := strings.Split(parsed[1], " ")
		scratchcard := strings.Split(parsed[2], " ")

		winningNums := intersect.Simple(winners, scratchcard)
		score = ScoreWinningScratchcard(winningNums)

		tally += score
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Final Tally: ", tally)

}
