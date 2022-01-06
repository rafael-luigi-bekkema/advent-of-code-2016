package day08

import (
	"log"
	"os"
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)




func TestDay8(t *testing.T) {
	result, _ := day8([]string{"rect 3x2", "rotate column x=1 by 1", "rotate row y=0 by 4",
		"rotate column x=1 by 1"}, 8, 3)
	TestEqual(t, 6, result)
	log.SetOutput(os.Stdout)
	result, out := day8(Lines(8), 50, 6)
	TestEqual(t, 106, result)
	expectDisp := "" +
		" ##  #### #    #### #     ##  #   #####  ##   ### \n" +
		"#  # #    #    #    #    #  # #   ##    #  # #    \n" +
		"#    ###  #    ###  #    #  #  # # ###  #    #    \n" +
		"#    #    #    #    #    #  #   #  #    #     ##  \n" +
		"#  # #    #    #    #    #  #   #  #    #  #    # \n" +
		" ##  #    #### #### ####  ##    #  #     ##  ###  \n"
	TestEqual(t, len(expectDisp), len(out))
	TestEqual(t, expectDisp, out, "CFLELOYFCS")
}
