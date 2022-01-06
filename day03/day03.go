package day03

import (
	"fmt"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func day3b(input [][3]int) int {
	var count int
	f := func(s1, s2, s3 int) {
		if s1+s2 > s3 && s2+s3 > s1 && s1+s3 > s2 {
			count++
		}
	}
	for i := 0; i < len(input); i += 3 {
		f(input[i][0], input[i+1][0], input[i+2][0])
		f(input[i][1], input[i+1][1], input[i+2][1])
		f(input[i][2], input[i+1][2], input[i+2][2])
	}
	return count
}

func day3a(input [][3]int) int {
	var count int
	for _, nums := range input {
		s1, s2, s3 := nums[0], nums[1], nums[2]
		if s1+s2 > s3 && s2+s3 > s1 && s1+s3 > s2 {
			count++
		}
	}
	return count
}

func day3input() (input [][3]int) {
	s, close := util.InputScanner(3)
	defer close()
	for s.Scan() {
		input = append(input, [3]int{})
		t := &input[len(input)-1]
		fmt.Sscanf(s.Text(), "%d %d %d", &t[0], &t[1], &t[2])
	}
	return
}
