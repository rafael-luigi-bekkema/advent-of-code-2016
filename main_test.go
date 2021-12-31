package main

import (
	"testing"
)

func TestDay3(t *testing.T) {
	TestEqual(t, 2, day3a([][3]int{{5, 6, 4}, {5, 10, 25}, {5, 5, 5}}))
	input := day3input()
	TestEqual(t, 982, day3a(input))
	TestEqual(t, 1826, day3b(input))
}

func TestDay2b(t *testing.T) {
	TestEqual(t, 0x5DB3, day2b([]string{"ULL", "RRDDD", "LURDL", "UUUUD"}))
	TestEqual(t, 0x3CC43, day2b(Lines(2)))
}

func TestDay2a(t *testing.T) {
	TestEqual(t, 1985, day2a([]string{"ULL", "RRDDD", "LURDL", "UUUUD"}))
	TestEqual(t, 19636, day2a(Lines(2)))
}

func TestDay1a(t *testing.T) {
	TestEqual(t, 5, day1a("R2, L3"))
	TestEqual(t, 2, day1a("R2, R2, R2"))
	TestEqual(t, 12, day1a("R5, L5, R5, R3"))
	TestEqual(t, 288, day1a(Input(1)))
}
func TestDay1b(t *testing.T) {
	TestEqual(t, 4, day1b("R8, R4, R4, R8"))
	TestEqual(t, 111, day1b(Input(1)))
}
