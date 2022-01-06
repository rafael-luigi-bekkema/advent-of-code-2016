package day12

import (
	"strings"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func day12b(input []string) int {
	return day12(input, true)
}

func day12a(input []string) int {
	return day12(input, false)
}

func day12(input []string, b bool) int {
	type CPU struct {
		regs    [4]int
		counter int
	}

	var cpu CPU
	if b {
		cpu.regs[2] = 1
	}

	set := func(s string, v int) {
		cpu.regs[s[0]-'a'] = v
	}

	get := func(s string) int {
		if '0' <= s[0] && s[0] <= '9' || s[0] == '-' {
			return Atoi(s)
		}
		return cpu.regs[s[0]-'a']
	}

	run := func(ins string) {
		cmd, rest, _ := strings.Cut(ins, " ")
		switch cmd {
		case "cpy":
			from, to, _ := strings.Cut(rest, " ")
			set(to, get(from))
		case "inc":
			set(rest, get(rest)+1)
		case "dec":
			set(rest, get(rest)-1)
		case "jnz":
			x, jmp, _ := strings.Cut(rest, " ")
			if get(x) != 0 {
				cpu.counter += get(jmp) - 1
			}
		default:
			panic("unknown cmd")
		}
	}

	for cpu.counter = 0; cpu.counter < len(input); cpu.counter++ {
		run(input[cpu.counter])
	}

	return cpu.regs[0] // register 'a'
}
