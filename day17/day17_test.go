package day17

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay17(t *testing.T) {
	TestEqual(t, "DDRRRD", day17a("ihgpwlah"))
	TestEqual(t, "DDUDRLRRUDRD", day17a("kglvqrro"))
	TestEqual(t, "DRURDRUDDLLDLUURRDULRLDUUDDDRR", day17a("ulqzkmiv"))
	TestEqual(t, "RDRRULDDDR", day17a("vkjiggvb"))

	TestEqual(t, 370, day17b("ihgpwlah"))
	TestEqual(t, 492, day17b("kglvqrro"))
	TestEqual(t, 830, day17b("ulqzkmiv"))
	TestEqual(t, 392, day17b("vkjiggvb"))
}
