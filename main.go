package main

import (
	"log"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/day11"
)

func main() {
	log.Println("Advent of Code 2016")
	// log.Println("Not much to see here. Run the tests:\ngo test -v ./...")

	day11.Day11b()
}

func init() {
	log.SetFlags(0)
}
