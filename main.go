package main

import (
	"log"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/day25"
)

func main() {
	log.Println("Advent of Code 2016")
	// log.Println("Not much to see here. Run the tests:\ngo test -v ./...")

	day25.Day25a()
}

func init() {
	log.SetFlags(0)
}
