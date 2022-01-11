package day25

import (
	"fmt"
	"strings"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func assembunny(cmds []string, aval int) bool {
	type CPU struct {
		regs    [4]int
		counter int
	}

	var cpu CPU
	cpu.regs[0] = aval

	set := func(s string, v int) {
		cpu.regs[s[0]-'a'] = v
	}

	get := func(s string) int {
		if '0' <= s[0] && s[0] <= '9' || s[0] == '-' {
			return Atoi(s)
		}
		return cpu.regs[s[0]-'a']
	}

	cmpins := func(inss ...string) bool {
		for i, ins := range inss {
			if cmds[cpu.counter+i] != ins {
				return false
			}
		}
		return true
	}

	var count, prev int
	var err bool

	run := func(ins string) {
		cmd, rest, _ := strings.Cut(ins, " ")
		switch cmd {
		case "cpy":
			if cmpins("cpy b c", "inc a", "dec c", "jnz c -2", "dec d", "jnz d -5") {
				set("a", get("d")*get("b"))
				set("c", 0)
				set("d", 0)
				cpu.counter += 5
				break
			}
			from, to, _ := strings.Cut(rest, " ")
			set(to, get(from))
		case "inc":
			if cmpins("inc a", "inc d", "jnz d -2") {
				set("a", get("a")+get("d")*-1)
				set("d", 0)
				cpu.counter += 2
				break
			}
			if cmpins("inc a", "dec d", "jnz d -2") {
				set("a", get("a")+get("d"))
				set("d", 0)
				cpu.counter += 2
				break
			}
			set(rest, get(rest)+1)
		case "dec":
			if cmpins("dec d", "inc c", "jnz d -2") {
				set("c", get("c")+get("d"))
				set("d", 0)
				cpu.counter += 2
				break
			}
			set(rest, get(rest)-1)
		case "jnz":
			x, jmp, _ := strings.Cut(rest, " ")
			if xval := get(x); xval != 0 {
				cpu.counter += get(jmp) - 1
			}
		case "out":
			val := get(rest)
			if count == 0 && val != 0 {
				err = true
				break
			}
			if count > 0 && (prev == 1 && val != 0 || prev == 0 && val != 1) {
				err = true
				break
			}
			prev = val
			count++
		case "tgl":
			jmp := get(rest)
			newi := cpu.counter + jmp
			if newi < 0 || newi > len(cmds)-1 {
				break
			}
			fmt.Printf("change ins at %d\n", newi)
			nins := cmds[newi]
			ncmd, nrest, _ := strings.Cut(nins, " ")
			switch ncmd {
			case "inc":
				cmds[newi] = fmt.Sprintf("dec %s", nrest)
			case "dec", "tgl":
				cmds[newi] = fmt.Sprintf("inc %s", nrest)

			case "jnz":
				cmds[newi] = fmt.Sprintf("cpy %s", nrest)
			case "cpy":
				cmds[newi] = fmt.Sprintf("jnz %s", nrest)
			}
		default:
			panic("unknown cmd: " + ins)
		}
	}

	for cpu.counter = 0; cpu.counter < len(cmds); cpu.counter++ {
		run(cmds[cpu.counter])
		if err {
			return false
		}
		if count > 20 {
			return true
		}
	}

	return false // register 'a'
}

func day25a(cmds []string) int {
	var i int
	for {
		if assembunny(cmds, i) {
			break
		}
		i++
	}
	return i
}

func Day25a() {
	fmt.Println("day 25a:", day25a(Lines(25)))
}
