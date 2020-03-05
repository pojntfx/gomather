package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type AddAndSubtractResults struct {
	x int64
	y int64
	e int64
}

var addResults = []AddAndSubtractResults{
	{0, 0, 0},
	{0, 1, 1},
	{1, 0, 1},
	{1, 1, 2},
	{-1, 1, 0},
	{1, -1, 0},
	{-1, -1, -2},
}
var subtractResults = []AddAndSubtractResults{
	{0, 0, 0},
	{1, 0, -1},
	{0, 1, 1},
	{-1, 0, 1},
	{-1, -1, 0},
	{1, -1, -2},
}

func TestAdd(t *testing.T) {
	for _, r := range addResults {
		assert.Equal(t, r.e, Add(r.x, r.y))
	}
}
func TestSubtract(t *testing.T) {
	for _, r := range subtractResults {
		assert.Equal(t, r.e, Subtract(r.x, r.y))
	}
}
