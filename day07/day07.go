package day07

import "strings"

func day7abba(input string) bool {
	abba := func(s string) bool {
		for i := 0; i < len(s)-3; i++ {
			if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
				return true
			}
		}
		return false
	}

	var addr, hyp string
	isAbba := false
	for {
		var ok bool
		addr, input, ok = strings.Cut(input, "[")
		if abba(addr) {
			isAbba = true
		}
		if !ok {
			break
		}
		hyp, input, ok = strings.Cut(input, "]")
		if abba(hyp) {
			// fmt.Println("hyper abba", hyp)
			return false
		}
	}
	return isAbba
}

func day7(input []string) (a int, b int) {
	for _, line := range input {
		if day7abba(line) {
			a++
		}
		if day7ssl(line) {
			b++
		}
	}
	return
}

func day7ssl(input string) bool {
	aba := func(s string) (a []string) {
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+2] && s[i] != s[i+1] {
				a = append(a, s[i:i+3])
			}
		}
		return
	}
	bab := func(hyp, aba string) bool {
		for i := 0; i < len(hyp)-2; i++ {
			if hyp[i] == aba[1] && hyp[i+1] == aba[0] && hyp[i+2] == aba[1] {
				return true
			}
		}
		return false
	}

	var addr, hyp string
	var hypers []string
	var abas []string
	for {
		var ok bool
		addr, input, ok = strings.Cut(input, "[")
		if a := aba(addr); a != nil {
			abas = append(abas, a...)
		}
		if !ok {
			break
		}
		hyp, input, _ = strings.Cut(input, "]")
		hypers = append(hypers, hyp)
	}
	for _, aba := range abas {
		for _, hyp := range hypers {
			if bab(hyp, aba) {
				return true
			}
		}
	}
	return false
}
