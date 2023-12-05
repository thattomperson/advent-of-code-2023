package main

import (
	"testing"

	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"
	"ttp.sh/advent-of-code/2023/util"
)

func TestFindFirst(t *testing.T) {
	assert.Equal(t, 1, FindFirst("onetwo", numbers))
	assert.Equal(t, 1, FindFirst("fivonetwo", numbers))
	assert.Equal(t, 1, FindFirst("oneight", numbers))
	assert.Equal(t, 2, FindFirst("two", numbers))
	assert.Equal(t, 7, FindFirst("seven", numbers))
	assert.Equal(t, 7, FindFirst("se7en", numbers))
	assert.Equal(t, 2, FindFirst("se2en", numbers))
}

func TestFindLast(t *testing.T) {
	assert.Equal(t, 2, FindLast("onetwo", numbers))
	assert.Equal(t, 2, FindLast("fivonetwo", numbers))
	assert.Equal(t, 8, FindLast("oneight", numbers))
	assert.Equal(t, 2, FindLast("two", numbers))
	assert.Equal(t, 7, FindLast("seven", numbers))
	assert.Equal(t, 7, FindLast("se7en", numbers))
	assert.Equal(t, 2, FindLast("se2en", numbers))
}

func TestRun(t *testing.T) {
	lines := []string{"one73", "2five2", "one2one", "oneight"}
	out := Run(util.MakeChan(lines))
	if out != 13+22+11+18 {
		t.FailNow()
	}
}
