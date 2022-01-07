package day20

import "fmt"

func day20(max int, rules []string) (int, int) {
	ips := make([]bool, max+1)
	for _, rule := range rules {
		var from, to int
		fmt.Sscanf(rule, "%d-%d", &from, &to)
		for i := from; i <= to; i++ {
			ips[i] = true
		}
	}
	count, first := 0, -1
	for i, ip := range ips {
		if !ip {
			if first == -1 {
				first = i
			}
			count++
		}
	}
	return first, count
}
