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

	switch os.Args[1] {
	case "toss":
		p.Event = pickr.EventToss
	case "roll":
		p.Event = pickr.EventRoll
	case "choose":
		p.Event = pickr.EventChoose
	default:
		fmt.Println("invalid event")
		printUsage()
		os.Exit(1)
	}

	seed := time.Now().UnixNano()
	p.Do(seed, os.Args[2:]...)
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
