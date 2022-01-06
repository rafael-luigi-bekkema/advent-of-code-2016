package day19

import (
	"container/ring"
)

func day19b3(nr int) int {
	cur := ring.New(nr)
	var ptarget *ring.Ring
	for i := 0; i < nr; i++ {
		cur.Value = i + 1
		if i == nr/2-1 {
			ptarget = cur
		}
		cur = cur.Next()
	}

	for size := nr; size > 1; size-- {
		ptarget.Unlink(1)
		ptarget = ptarget.Move(size % 2)
	}
	return ptarget.Value.(int)
}

func day19a(nr int) int {
	first := 1
	for step := 2; nr > 1; nr, step = nr/2, step*2 {
		first += step * (nr % 2)
	}
	return first
}
