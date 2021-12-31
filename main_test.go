package main

import (
	"log"
	"os"
	"testing"
)

func TestDay8(t *testing.T) {
	result, _ := day8([]string{"rect 3x2", "rotate column x=1 by 1", "rotate row y=0 by 4",
		"rotate column x=1 by 1"}, 8, 3)
	TestEqual(t, 6, result)
	log.SetOutput(os.Stdout)
	result, out := day8(Lines(8), 50, 6)
	TestEqual(t, 106, result)
	expectDisp := "" +
		" ##  #### #    #### #     ##  #   #####  ##   ### \n" +
		"#  # #    #    #    #    #  # #   ##    #  # #    \n" +
		"#    ###  #    ###  #    #  #  # # ###  #    #    \n" +
		"#    #    #    #    #    #  #   #  #    #     ##  \n" +
		"#  # #    #    #    #    #  #   #  #    #  #    # \n" +
		" ##  #    #### #### ####  ##    #  #     ##  ###  \n"
	TestEqual(t, len(expectDisp), len(out))
	TestEqual(t, expectDisp, out)
}

func TestDay7(t *testing.T) {
	TestEqual(t, true, day7abba("abba[mnop]qrst"))
	TestEqual(t, false, day7abba("abcd[bddb]xyyx"))
	testA, _ := day7([]string{"abba[mnop]qrst", "abcd[bddb]xyyx", "aaaa[qwer]tyui", "ioxxoj[asdfgh]zxcvbn"})
	TestEqual(t, 2, testA)
	resultA, resultB := day7(Lines(7))
	TestEqual(t, 118, resultA)

	TestEqual(t, true, day7ssl("aba[bab]xyz"))
	TestEqual(t, false, day7ssl("xyx[xyx]xyx"))
	TestEqual(t, true, day7ssl("aaa[kek]eke"))
	TestEqual(t, true, day7ssl("zazbz[bzb]cdb"))
	TestEqual(t, 260, resultB)
}

func TestDay6(t *testing.T) {
	input := []string{"eedadn", "drvtee", "eandsr", "raavrd", "atevrs", "tsrnev", "sdttsa", "rasrtv",
		"nssdts", "ntnada", "svetve", "tesnvt", "vntsnd", "vrdear", "dvrsen", "enarar"}
	file := Lines(6)
	TestEqual(t, "easter", day6a(input))
	TestEqual(t, "tsreykjj", day6a(file))

	TestEqual(t, "advent", day6b(input))
	TestEqual(t, "hnfbujie", day6b(file))
}

func TestDay5(t *testing.T) {
	TestEqual(t, "18f47a30", day5a("abc"))
	TestEqual(t, "2414bc77", day5a("wtnhxymk"))

	TestEqual(t, "05ace8e3", day5b("abc"))
	TestEqual(t, "437e60fc", day5b("wtnhxymk"))
}

func TestDay4a(t *testing.T) {
	result, _ := day4a([]string{
		"aaaaa-bbb-z-y-x-123[abxyz]",
		"a-b-c-d-e-f-g-h-987[abcde]",
		"not-a-real-room-404[oarel]",
		"totally-real-room-200[decoy]"})
	TestEqual(t, 1515, result)
	result, resultID := day4a(Lines(4))
	TestEqual(t, 245102, result)
	TestEqual(t, "very encrypted name", day4b("qzmt-zixmtkozy-ivhz", 343))
	TestEqual(t, 324, resultID)
}

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
