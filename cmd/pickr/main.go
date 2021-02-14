package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/thealamu/pickr"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	p := &pickr.Pickr{Out: os.Stdout}
	now := time.Now().UnixNano()

	switch os.Args[1] {
	case "toss":
		p.Event = pickr.EventToss
		fmt.Println("Tossing a coin")
		p.Do(now, os.Args[2:]...)

	case "roll":
		p.Event = pickr.EventRoll
		fmt.Println("Rolling a die")
		p.Do(now, os.Args[2:]...)

	case "choose":
		p.Event = pickr.EventChoose
		fmt.Println("Choosing from list")
		p.Do(now, os.Args[2:]...)

	default:
		fmt.Println("invalid event")
		printUsage()
		os.Exit(1)
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
