package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	text, _ := os.ReadFile("./text.txt")

	text_lines := strings.Split(string(text), "\n")

	number_words := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var sum int = 0

	for _, line := range text_lines { // Lines

		if len(line) == 0 {
			continue
		}

		// fmt.Println(line)
		// fmt.Println(index)
		var first_number, last_number string = "", ""
		var first_number_pop bool = false

		for index_inner, number_word := range number_words {
			// fmt.Println(line, number_word, strconv.Itoa(index_inner+1))
			line = strings.Replace(line, number_word, number_word+strconv.Itoa(index_inner+1)+number_word, -1)
		}
		// fmt.Println(line)

		for _, rune := range line { // Characters
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

	fmt.Println("Part 2: ", sum)
}
