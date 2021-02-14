package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "toss":
		fmt.Println("Tossing a coin")
	case "roll":
		fmt.Println("Rolling a die")
	case "choose":
		fmt.Println("Choosing from list")
	}
}

func printUsage() {
	fmt.Println(strings.TrimSpace(`
Usage: pickr event
Make random choices. Toss a coin, roll a die, choose from a list of args.

Available events:
	toss		toss a coin
	roll n		roll a n-sided die
	choose ARGS...	choose from a list of args

	`))
}
