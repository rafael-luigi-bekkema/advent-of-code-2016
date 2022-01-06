package day09

import (
	"fmt"
	"strings"
)

func day9decompressB(input string) int {
	return day9decompress(input, true, 0)
}

func day9decompressA(input string) int {
	return day9decompress(input, false, 0)
}

func day9decompress(input string, recursive bool, depth int) int {
	var chunk string
	var ok bool
	var length, n int
	var count int
	for {
		chunk, input, ok = strings.Cut(input, "(")
		count += len(chunk)
		if !ok {
			break
		}
		chunk, input, ok = strings.Cut(input, ")")
		fmt.Sscanf(chunk, "%dx%d", &length, &n)

		chunk = input[:length]
		chunkLen := len(chunk)
		if recursive {
			chunkLen = day9decompress(chunk, true, depth+1)
		}
		count += n * chunkLen
		input = input[length:]
	}
	return count
}
