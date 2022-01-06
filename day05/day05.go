package day05

import (
	"crypto/md5"
	"fmt"
)

func day5b(doorID string) string {
	i := -1
	var found int
	pass := [8]byte{'_', '_', '_', '_', '_', '_', '_', '_'}
	for {
		i++
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(doorID, i))))
		if hash[:5] != "00000" {
			continue
		}
		pos := int(hash[5] - '0')
		if !(0 <= pos && pos <= 7) || pass[pos] != '_' {
			continue
		}

		pass[pos] = hash[6]
		found++

		fmt.Printf("\rfound %s %8d %s", string(pass[:]), i, hash)
		if found == 8 {
			fmt.Print("\n\n")
			break
		}
	}
	return string(pass[:])
}

func day5a(doorID string) string {
	i := -1
	pass := make([]byte, 0, 8)
	for {
		i++
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(doorID, i))))
		if hash[:5] != "00000" {
			continue
		}
		pass = append(pass, hash[5])
		if len(pass) == 8 {
			break
		}
	}
	return string(pass)
}
