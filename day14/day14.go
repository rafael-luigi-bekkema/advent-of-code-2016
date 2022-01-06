package day14

import (
	"crypto/md5"
	"fmt"
	"sort"
)


func day14b(salt string, nth int) int {
	return day14(salt, nth, true)
}

func day14a(salt string, nth int) int {
	return day14(salt, nth, false)
}

func day14(salt string, nth int, b bool) int {
	stretch := func(hash []byte) []byte {
		for i := 0; i < 2016; i++ {
			hash = []byte(fmt.Sprintf("%x", md5.Sum(hash)))
		}
		return hash
	}

	idx := 0
	type match struct {
		char byte
		hash []byte
		idx  int
	}
	var matches []match
	var keys []match
	for {
		hash := []byte(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", salt, idx)))))
		if b {
			hash = stretch(hash)
		}

		for i, c := range hash[:len(hash)-4] {
			if !(c == hash[i+1] && c == hash[i+2] && c == hash[i+3] && c == hash[i+4]) {
				continue
			}
		matcher:
			for mi := len(matches) - 1; mi >= 0; mi-- {
				match := &matches[mi]
				if match.idx+1000 < idx {
					matches[mi], matches[len(matches)-1] = matches[len(matches)-1], matches[mi]
					matches = matches[:len(matches)-1]
					continue
				}
				if match.char == c {
					keys = append(keys, *match)
					matches[mi], matches[len(matches)-1] = matches[len(matches)-1], matches[mi]
					matches = matches[:len(matches)-1]
					sort.Slice(keys, func(i, j int) bool {
						return keys[i].idx < keys[j].idx
					})
					if len(keys) >= nth {
						maxIdx := keys[nth-1].idx
						for _, match := range matches {
							if match.idx < maxIdx {
								continue matcher
							}
						}
						return keys[nth-1].idx
					}
				}
			}
			break
		}

		for i, c := range hash[:len(hash)-2] {
			if c == hash[i+1] && c == hash[i+2] {
				matches = append(matches, match{c, hash, idx})
				break
			}
		}

		idx++
		if idx > 2372900 {
			break
		}
	}
	return 0
}

