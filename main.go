package main

import (
	"advent-of-code-2024/day01"
	"advent-of-code-2024/day02"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	dayAndPart := flag.Arg(0)

	switch dayAndPart {
	case "0101":
		day01.Part1()
	case "0102":
		day01.Part2()
	case "0201":
		day02.Part1()
	case "":
		fmt.Println("Please choose a day and part to run, in the format DDPP.")
	default:
		fmt.Println("Please choose a day and part to run, in the format DDPP.")
	}
}
