package main

import (
	"log"
	"os"
	"testing"
)

func TestDay16(t *testing.T) {
	TestEqual(t, "01100", day16a(20, "10000"))
	TestEqual(t, "01110011101111011", day16a(272, "11110010111001001"))
	TestEqual(t, "11001111011000111", day16a(35651584, "11110010111001001"))
}

func TestDay15(t *testing.T) {
	TestEqual(t, 5, day15a([]string{
		"Disc #1 has 5 positions; at time=0, it is at position 4.",
		"Disc #2 has 2 positions; at time=0, it is at position 1.",
	}))
	TestEqual(t, 122318, day15a(Lines(15)))
	TestEqual(t, 3208583, day15b(Lines(15)))
}

func TestDay14(t *testing.T) {
	TestEqual(t, 22728, day14a("abc", 64))
	TestEqual(t, 15168, day14a("qzyelonm", 64))
	TestEqual(t, 22551, day14b("abc", 64))
	TestEqual(t, 20864, day14b("qzyelonm", 64))
}

func TestDay13(t *testing.T) {
	testa, _ := day13a(10, Coord{7, 4})
	TestEqual(t, 11, testa)
	a, b := day13a(1358, Coord{31, 39})
	TestEqual(t, 96, a)
	TestEqual(t, 141, b)
}

func TestDay12(t *testing.T) {
	TestEqual(t, 42, day12a([]string{"cpy 41 a",
		"inc a",
		"inc a",
		"dec a",
		"jnz a 2",
		"dec a",
	}))
	TestEqual(t, 6, day12a([]string{"cpy 5 a",
		"jnz a 2",
		"inc a",
		"inc a",
	}))
	TestEqual(t, 317993, day12a(Lines(12)))
	TestEqual(t, 9227647, day12b(Lines(12)))
}

func TestDay11(t *testing.T) {
	// TestEqual(t, 11, day11a(
	// 	[4][]string{
	// 		{"HM", "LM"},
	// 		{"HG"},
	// 		{"LG"},
	// 		nil,
	// 	},
	// ))

	/*
		F4 .  .  .  .  .  .  .  .  .  .  .  .
		F3 .  .  .  .  .  .  .  TM .  .  .  .
		F2 .  .  .  .  .  .  TG . RG RM CG CM
		F1 E  SG  SM PG  PM  .  .  .  .  .  .
	*/
	TestEqual(t, -1, day11a(
		[4][]string{
			{"SM", "PM", "SG", "PG"},
			{"RM", "CM", "TG", "RG", "CG"},
			{"TM"},
			nil,
		},
	))
}

func TestDay10(t *testing.T) {
	TestEqual(t, 2, day10([]string{
		"value 5 goes to bot 2",
		"bot 2 gives low to bot 1 and high to bot 0",
		"value 3 goes to bot 1",
		"bot 1 gives low to output 1 and high to bot 0",
		"bot 0 gives low to output 2 and high to output 0",
		"value 2 goes to bot 2",
	}, 2, 5, false))
	file := Lines(10)
	TestEqual(t, 47, day10(file, 17, 61, false))
	TestEqual(t, 2666, day10(file, 17, 61, true))
}

func TestDay9(t *testing.T) {
	TestEqual(t, 6, day9decompressA("ADVENT"))
	TestEqual(t, 7, day9decompressA("A(1x5)BC"))
	TestEqual(t, 9, day9decompressA("(3x3)XYZ"))
	TestEqual(t, 11, day9decompressA("A(2x2)BCD(2x2)EFG"))
	TestEqual(t, 6, day9decompressA("(6x1)(1x3)A"))
	TestEqual(t, 18, day9decompressA("X(8x2)(3x3)ABCY"))
	file := Input(9)
	TestEqual(t, 115118, day9decompressA(file))

	TestEqual(t, 9, day9decompressB("(3x3)XYZ"))
	TestEqual(t, 20, day9decompressB("X(8x2)(3x3)ABCY"))
	TestEqual(t, 241920, day9decompressB("(27x12)(20x12)(13x14)(7x10)(1x12)A"))
	TestEqual(t, 445, day9decompressB("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"))

	TestEqual(t, 11107527530, day9decompressB(file))
}

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
	TestEqual(t, expectDisp, out, "CFLELOYFCS")
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
