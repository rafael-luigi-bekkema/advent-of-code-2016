package day04

import (
	"fmt"
	"sort"
	"strings"
)

func day4b(input string, sid int) string {
	chars := []byte(strings.ReplaceAll(input, "-", " "))
	for i := 0; i < sid; i++ {
		for j, c := range chars {
			if c == ' ' {
				continue
			}
			if c == 'z' {
				chars[j] = 'a'
				continue
			}
			chars[j]++
		}
	}
	return string(chars)
}

func day4a(input []string) (int, int) {
	checksum := func(inp string) string {
		cc := map[rune]int{}
		var chars []rune
		for _, c := range inp {
			if c == '-' {
				continue
			}
			if _, ok := cc[c]; !ok {
				chars = append(chars, c)
			}
			cc[c]++
		}
		sort.Slice(chars, func(i, j int) bool {
			if cc[chars[i]] == cc[chars[j]] {
				return chars[i] < chars[j]
			}
			return cc[chars[i]] > cc[chars[j]]
		})
		return string(chars[:5])
	}

	var total int
	var partBid int
	for _, line := range input {
		lidx := strings.LastIndex(line, "-")
		csum := checksum(line[:lidx])

		var sid int
		var sum string
		fmt.Sscanf(line[lidx:], "-%d[%s]", &sid, &sum)
		decoded := day4b(line[:lidx], sid)
		if decoded == "northpole object storage" {
			partBid = sid
		}
		if sum[:len(sum)-1] == csum {
			total += sid
		}
	}
	return total, partBid
}
