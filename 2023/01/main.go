package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "unicode"
)


func FindAllNumbers(input string) []int {
	var output []int
  runeInput := []rune(input)

  if strings.ContainsAny(input, "1234567890") {
		for _, x := range runeInput {
			if unicode.IsDigit(x) {
				// Convert from ASCII table entry into actual digit
				output = append(output, (int(x) - 48))
		  }
 		}
  } else {
	  output = append(output, 0) 
  }
  return output
}


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

		totalCount := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        nums := FindAllNumbers(scanner.Text())
				fmt.Println(scanner.Text(), nums)
				switch {
					case len(nums) == 0:
						continue
					case len(nums) == 1:
						totalCount = totalCount + (nums[0] * 10) + (nums[0])
					case len(nums) > 1:
						totalCount = totalCount + (nums[0] * 10) + (nums[len(nums)-1])
				}

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	  fmt.Println("Final Total: ", totalCount)
}
