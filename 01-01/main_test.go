package main

import (
	"testing"

	"ttp.sh/advent-of-code/2023/util"
)

func Test(t *testing.T) {
	lines := []string{"173", "252", "1"}
	out := Run(util.MakeChan(lines))
	if out != 13+22+11 {
		t.FailNow()
	}
}
