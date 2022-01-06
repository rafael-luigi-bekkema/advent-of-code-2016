package day16

import "fmt"


func day16a(length int, initial string) string {
	flip := func(i byte) byte {
		if i == '0' {
			return '1'
		}
		return '0'
	}
	a := initial
	for len(a) < length {
		b := []byte(a)
		for i, j := 0, len(b)-1; i <= j; i, j = i+1, j-1 {
			b[i], b[j] = flip(b[j]), flip(b[i])
		}
		a = fmt.Sprintf("%s0%s", a, b)
	}
	cs := a[:length]
	for {
		ncs := make([]byte, len(cs)/2)
		for i := 0; i < len(cs); i += 2 {
			if cs[i] == cs[i+1] {
				ncs[i/2] = '1'
			} else {
				ncs[i/2] = '0'
			}
		}
		cs = string(ncs)
		if len(cs)%2 != 0 {
			break
		}
	}
	return cs
}
