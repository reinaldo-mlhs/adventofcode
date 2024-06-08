package main

import (
	"flag"
	"fmt"
)

func main() {
	year_flag := flag.String("year", "2023", "a string")
	day_flag := flag.String("day", "1", "a string")
	part_flag := flag.String("part", "1", "a string")

	flag.Parse()

	fmt.Println("Year: ", *year_flag)
	fmt.Println("Day: ", *day_flag)
	fmt.Println("Part: ", *part_flag)

}
