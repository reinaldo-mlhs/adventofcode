package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	text, _ := os.ReadFile("./text.txt")

	text_lines := strings.Split(string(text), "\n")

	var sum int = 0

	for _, line := range text_lines {

		if len(line) == 0 {
			continue
		}

		// fmt.Println(line)
		// fmt.Println(index)
		var first_number, last_number string = "", ""
		var first_number_pop bool = false

		for _, rune := range line {
			char, _ := strconv.Unquote(strconv.QuoteRune(rune))
			_, error := strconv.Atoi(string(char))

			if error != nil {
				continue
			}

			if !first_number_pop {
				first_number = char
				first_number_pop = true
			}

			last_number = char
		}

		// fmt.Println(first_number, last_number)
		var final_number string = first_number + last_number
		// fmt.Println(final_number)
		number, error := strconv.Atoi(final_number)

		if error != nil {
			panic(error)
		}

		sum = sum + number
		// fmt.Println(sum)
	}

	fmt.Println("Part 1: ", sum)
}
