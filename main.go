package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	dayAndPart := flag.Arg(0)

	switch dayAndPart {
	case "0101":
		fmt.Printf(
			"day is %s\n",
			dayAndPart,
		)
	case "":
		fmt.Println("Please choose a day and part to run, in the format DDPP.")
	default:
		fmt.Println("Please choose a day and part to run, in the format DDPP.")
	}
}
