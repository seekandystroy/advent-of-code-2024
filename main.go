package main

import (
	"advent-of-code-2024/day01"
	"advent-of-code-2024/day02"
	"advent-of-code-2024/day03"
	"advent-of-code-2024/day04"
	"advent-of-code-2024/day05"
	"advent-of-code-2024/day06"
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
	case "0202":
		day02.Part2()
	case "0301":
		day03.Part1()
	case "0302":
		day03.Part2()
	case "0401":
		day04.Part1()
	case "0402":
		day04.Part2()
	case "0501":
		day05.Part1()
	case "0502":
		day05.Part2()
	case "0601":
		day06.Part1()
	case "0602":
		day06.Part2()
	case "":
		fmt.Println("Please choose a day and part to run, in the format DDPP.")
	default:
		fmt.Println("Please choose a day and part to run, in the format DDPP.")
	}
}
