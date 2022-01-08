package day23

import (
	"fmt"
	"strings"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func assembunny(instructions []string, initA bool) int {
	type CPU struct {
		regs    [4]int
		counter int
	}

	var cpu CPU
	if initA {
		cpu.regs[0] = 7
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
		case "tgl":
			jmp := get(rest)
			newi := cpu.counter + jmp
			if newi < 0 || newi > len(instructions)-1 {
				break
			}
			nins := instructions[newi]
			ncmd, nrest, _ := strings.Cut(nins, " ")
			switch ncmd {
			case "inc":
				instructions[newi] = fmt.Sprintf("dec %s", nrest)
			case "dec", "tgl":
				instructions[newi] = fmt.Sprintf("inc %s", nrest)

			case "jnz":
				instructions[newi] = fmt.Sprintf("cpy %s", nrest)
			case "cpy":
				instructions[newi] = fmt.Sprintf("jnz %s", nrest)
			}
		default:
			panic("unknown cmd: " + ins)
		}
	}

	for cpu.counter = 0; cpu.counter < len(instructions); cpu.counter++ {
		run(instructions[cpu.counter])
	}

	return cpu.regs[0] // register 'a'
}
