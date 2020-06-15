package list

import (
	"testing"
)

func TestInit(t *testing.T) {
	l := New()
	checkListPointers(t, l, []*Node{})
	if l.Head() != nil {
		t.Errorf("Head of empty list should be nil but was not")
	}
	if l.Tail() != nil {
		t.Errorf("Tail of empty list should be nil but was not")
	}
}

func TestLazyInit(t *testing.T) {
	l := List{}
	checkListPointers(t, &l, []*Node{})
	l.lazyInit()
	checkListPointers(t, &l, []*Node{})
}

func TestAdd(t *testing.T) {
	l := New()
	n1 := l.AddToBack("John")
	checkListPointers(t, l, []*Node{n1})
	n2 := l.AddToBack("Ashling")
	checkListPointers(t, l, []*Node{n1, n2})
	n3 := l.AddToFront("Finn")
	checkListPointers(t, l, []*Node{n3, n1, n2})
	n4 := l.AddBefore("Anthony", n2)
	checkListPointers(t, l, []*Node{n3, n1, n4, n2})
	n5 := l.AddAfter("Daniella", n1)
	checkListPointers(t, l, []*Node{n3, n1, n5, n4, n2})
}

func TestDelete(t *testing.T) {
	l := New()
	n1 := l.AddToBack("John")
	n2 := l.AddToBack("Ashling")
	n3 := l.AddToBack("Finn")
	data := l.Delete(n3)
	if d := data.(string); d != "Finn" {
		t.Errorf("Expected Finn but got %s", d)
	}
	checkListPointers(t, l, []*Node{n1, n2})
}

func TestIteration(t *testing.T) {
	l := New()
	l.AddToBack(1)
	l.AddToBack(1)
	l.AddToBack(1)

	// test forward iteration
	var i int
	var sum int
	for n := l.Head(); n != nil; n = n.Next() {
		if v, ok := n.Data.(int); ok {
			sum += v
		}
		i++
		if i > 3 {
			t.Errorf("iteration should end after 3 nodes")
			return
		}
	}

	if sum != 3 {
		t.Errorf("sum over 3 values should equal 3, but was %d", sum)
	}

	// test reverse iteration
	sum = 0
	i = 0
	for n := l.Tail(); n != nil; n = n.Prev() {
		if v, ok := n.Data.(int); ok {
			sum += v
		}
		i++
		if i > 3 {
			t.Errorf("iteration should end after 3 nodes")
			return
		}
	}

	if sum != 3 {
		t.Errorf("sum over 3 values should equal 3, but was %d", sum)
	}
}

func checkListPointers(t *testing.T, l *List, nodes []*Node) {

	// check that the size of the list is correct
	if !checkListSize(t, l, nodes) {
		return
	}

	// check that the list contains the right pointers based
	// on the nodes provided
	root := &l.root
	for i, n := range nodes {
		expPrev := root         // for checking the internal previous which will be the root for the first node
		ExpPrev := (*Node)(nil) // for checking the external previous which will be nil for the first node
		if i > 0 {
			expPrev = nodes[i-1]
			ExpPrev = expPrev
		}
		if expPrev != n.prev {
			t.Errorf("For node %d, the previous was expected to be %p, but was %p", i, expPrev, n.prev)
		}
		if ExpPrev != n.Prev() {
			t.Errorf("For node %d, the previous was expected to be %p, but was %p", i, expPrev, n.prev)
		}

		expNext := root         // for checking the internal previous which will be the root for the last node
		ExpNext := (*Node)(nil) // for checking the external previous which will be nil for the last node
		if i < len(nodes)-1 {
			expNext = nodes[i+1]
			ExpNext = expNext
		}
		if expNext != n.next {
			t.Errorf("For node %d, the next was expected to be %p, but was %p", i, expNext, n.next)
		}
		if ExpNext != n.Next() {
			t.Errorf("For node %d, the next was expected to be %p, but was %p", i, expNext, n.next)
		}

	}

}

func checkListSize(t *testing.T, l *List, nodes []*Node) bool {
	if l.Size != len(nodes) {
		t.Errorf("Size of list was %d but %d nodes were expected", l.Size, len(nodes))
		return false
	}

	root := &l.root
	if l.Size == 0 {
		if l.root.next != nil && l.root.next != root || l.root.prev != nil && l.root.prev != root {
			t.Errorf("0 length list must be zero values or circled back on itself")
		}
	}
	return true
}
