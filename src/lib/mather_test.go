// +build unit

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
		a := Add(r.x, r.y)
		assert.Equal(t, r.e, a, "%v + %v should be %v, but was %v", r.x, r.y, r.e, a)
	}
}
func TestSubtract(t *testing.T) {
	for _, r := range subtractResults {
		a := Subtract(r.x, r.y)
		assert.Equal(t, r.e, a, "%v - %v should be %v, but was %v", r.y, r.x, r.e, a)
	}
}
