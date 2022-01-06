package day10

import (
	"fmt"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func day10(input []string, checklow, checkhigh int, b bool) int {
	type Ins struct {
		frombot, tolow, tohigh int
		towhatlow, towhathigh  string
	}
	var instructions []Ins

	bots := map[int][]int{}
	for i := len(input) - 1; i >= 0; i-- {
		if input[i][:5] != "value" {
			instructions = append(instructions, Ins{})
			ins := &instructions[len(instructions)-1]
			Must(fmt.Sscanf(input[i], "bot %d gives low to %s %d and high to %s %d",
				&ins.frombot, &ins.towhatlow, &ins.tolow, &ins.towhathigh, &ins.tohigh))
			continue
		}
		var val, botnum int
		fmt.Sscanf(input[i], "value %d goes to bot %d", &val, &botnum)
		bots[botnum] = append(bots[botnum], val)

		input[i], input[len(input)-1] = input[len(input)-1], input[i]
		input = input[:len(input)-1]
	}

	outputs := map[int]int{}
	for i := len(instructions) - 1; i >= 0; i-- {
		ins := &instructions[i]

		if len(bots[ins.frombot]) < 2 {
			continue
		}

		low, high := MinMax(bots[ins.frombot])
		if !b && low == checklow && high == checkhigh {
			return ins.frombot
		}

		bots[ins.frombot] = nil
		if ins.towhatlow == "bot" {
			bots[ins.tolow] = append(bots[ins.tolow], low)
		} else {
			outputs[ins.tolow] = low
		}
		if ins.towhathigh == "bot" {
			bots[ins.tohigh] = append(bots[ins.tohigh], high)
		} else {
			outputs[ins.tohigh] = high
		}

		instructions[i], instructions[len(instructions)-1] = instructions[len(instructions)-1], instructions[i]
		instructions = instructions[:len(instructions)-1]

		i = len(instructions)
	}
	return outputs[0] * outputs[1] * outputs[2]
}
