package list

import (
	"testing"
)

func Test_Ring(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()

	var space [10]NodeX
	for i := 0; i < 10; i++ {
		space[i].Val = i
	}

	var ring Ring
	ring.Initialize()

	for i := 4; i >= 0; i-- {
		ring.InsertHead(&space[i])
	}
	for i := 5; i < 10; i++ {
		ring.InsertTail(&space[i])
	}

	Release(&space[8])
	ring.InsertTail(&space[8])
	Release(&space[9])
	ring.InsertHead(&space[9])
	Release(&space[8])
	ring.InsertTail(&space[8])
	Release(&space[0])
	ring.InsertTail(&space[0])

	//9 1 2 3 4 5 6 7 8 0

	var node = ring.PopHead()
	if node == nil || node.Val != 9 {
		t.Fail()
	}
	node = ring.PopTail()
	if node == nil || node.Val != 0 {
		t.Fail()
	}
	for i := 1; i < 9; i++ {
		node = ring.PopHead()
		if node == nil || node.Val != i {
			t.Fail()
		}
	}
	if !ring.IsEmpty() {
		t.Fail()
	}
}
