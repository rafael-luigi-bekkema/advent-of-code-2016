package util

import (
	"bufio"
	"constraints"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

var inputDir = "."

func init() {
	if idir, err := filepath.Abs(inputDir); err == nil {
		inputDir = idir
	}
}

func inputPath(day int, suffix ...string) string {
	inputDir := inputDir
	dirName := fmt.Sprintf("day%02d", day)
	if filepath.Base(inputDir) != dirName {
		inputDir = filepath.Join(inputDir, dirName)
	}

	var s string
	if len(suffix) > 0 {
		s = strings.Join(suffix, "")
	}
	return Must(filepath.Abs(fmt.Sprintf("%s/day%02d%s.txt", inputDir, day, s)))
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Atoi(input string) int {
	i := Must(strconv.Atoi(input))
	return i
}

func Abs[T constraints.Signed | constraints.Float](n T) T {
	if n < 0 {
		return n * -1
	}
	return n
}

func InputReader(day int, suffix ...string) io.ReadSeekCloser {
	return Must(os.Open(inputPath(day, suffix...)))
}

func InputScanner(day int, suffix ...string) (s *bufio.Scanner, close func() error) {
	f := InputReader(day, suffix...)
	s = bufio.NewScanner(f)
	return s, f.Close
}

func Lines(day int, suffix ...string) []string {
	b, close := InputScanner(day, suffix...)
	defer close()
	var lines []string
	for b.Scan() {
		lines = append(lines, b.Text())
	}
	return lines
}

func Input(day int, suffix ...string) string {
	data := Must(os.ReadFile(inputPath(day, suffix...)))
	return strings.TrimSpace(string(data))
}

func In[T comparable](item T, items ...T) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func TestEqual[T comparable](t *testing.T, expect, result T, message ...string) {
	t.Helper()
	var msg string
	if len(message) > 0 {
		msg = message[0]
	}
	t.Run(msg, func(t *testing.T) {
		t.Helper()
		if result != expect {
			t.Fatalf("expected %v, got %v", expect, result)
		}
	})
}

func MinMax[T constraints.Integer | constraints.Float](values []T) (min T, max T) {
	for i, v := range values {
		if i == 0 || v < min {
			min = v
		}
		if i == 0 || v > max {
			max = v
		}
	}
	return min, max
}

func Remove[T comparable](list []T, remove ...T) []T {
	var newlist []T
outer:
	for i := range list {
		for j := range remove {
			if list[i] == remove[j] {
				continue outer
			}
		}
		newlist = append(newlist, list[i])
	}
	return newlist
}

func MD5(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}
