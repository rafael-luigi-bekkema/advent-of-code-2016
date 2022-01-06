package day14

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay14(t *testing.T) {
	TestEqual(t, 22728, day14a("abc", 64))
	TestEqual(t, 15168, day14a("qzyelonm", 64))
	TestEqual(t, 22551, day14b("abc", 64))
	TestEqual(t, 20864, day14b("qzyelonm", 64))
}
