package flow

import (
	"testing"
)

//seperate & search ?

func Test_Dinic(t *testing.T) {
	var matrix = [][]uint{
		{0, 8, 0, 4, 0, 0, 11, 0, 0},
		{8, 0, 7, 0, 2, 0, 0, 0, 4},
		{0, 7, 0, 0, 0, 9, 0, 0, 14},
		{4, 0, 0, 0, 0, 0, 8, 0, 0},
		{0, 2, 0, 0, 0, 0, 7, 6, 0},
		{0, 0, 9, 0, 0, 0, 0, 0, 10},
		{11, 0, 0, 8, 7, 0, 0, 1, 0},
		{0, 0, 0, 0, 6, 0, 1, 0, 2},
		{0, 4, 14, 0, 0, 10, 0, 2, 0}}
	var ret = Dinic(matrix)
	if ret != 12 {
		t.Fail()
	}

	matrix = [][]uint{
		{0, 16, 13, 0, 0, 0},
		{0, 0, 0, 12, 0, 0},
		{0, 4, 0, 0, 14, 0},
		{0, 0, 0, 0, 0, 20},
		{0, 0, 0, 7, 0, 4},
		{0, 0, 0, 0, 0, 0}}
	ret = Dinic(matrix)
	if ret != 23 {
		t.Fail()
	}

	matrix = [][]uint{
		{0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0}}
	ret = Dinic(matrix)
	if ret != 2 {
		t.Fail()
	}
}
