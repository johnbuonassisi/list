package list

import (
	"testing"
)

func TestAdd(t *testing.T) {
	list := List{}
	list.Add("John")
	head := list.Head()
	if head == nil {
		t.Fatal("the head element was nil")
	}
	tail := list.Tail()
	if tail == nil {
		t.Fatal("the tail element was nil")
	}
	if tail != head {
		t.Fatalf("the head and tail elements are not the same: %v, %v", head, tail)
	}
	if head.element != "John" {
		t.Fatalf("the head element is not John: %v", head)
	}
	list.Add("Ashling")
	head = list.Head()
	tail = list.Tail()
	if head.element != "John" {
		t.Fatalf("the head element is not John: %v", head)
	}
	if tail.element != "Ashling" {
		t.Fatalf("the tail element is not Ashling: %v", head)
	}
	t.Logf("ran test %s", "John")
}

func TestIterate(t *testing.T) {
	list := NewWithElements("John", "Ashling", "Finn")
	i := 0
	for node := list.Head(); node != nil; node = node.Next() {
		if i == 0 && node.Element() != "John" {
			t.Fatalf("element number, %d, should be John", i)
		} else if i == 1 && node.Element() != "Ashling" {
			t.Fatalf("element number, %d, should be Ashling", i)
		} else if i == 2 && node.Element() != "Finn" {
			t.Fatalf("element number, %d, should be Finn", i)
		}
		i = i + 1
	}
}

func TestRevIterate(t *testing.T) {
	list := NewWithElements("John", "Ashling", "Finn")
	i := 2
	for node := list.Tail(); node != nil; node = node.Prev() {
		if i == 0 && node.Element() != "John" {
			t.Fatalf("element number, %d, should be John", i)
		} else if i == 1 && node.Element() != "Ashling" {
			t.Fatalf("element number, %d, should be Ashling", i)
		} else if i == 2 && node.Element() != "Finn" {
			t.Fatalf("element number, %d, should be Finn", i)
		}
		i = i - 1
	}
}

func TestIterateAfterInsert(t *testing.T) {
	list := NewWithElements("John", "Ashling", "Finn")
	list.Insert("Daniella", 0)
	list.Insert("Anthony", 4)
	list.Insert("Peter", 2)
	// result is Daniella, John, Peter, Ashling, Finn, Anthony
	result := map[int]string{0: "Daniella", 1: "John", 2: "Peter", 3: "Ashling", 4: "Finn", 5: "Anthony"}
	var i int
	for node := list.Head(); node != nil; node = node.Next() {
		if result[i] != node.Element() {
			t.Fatalf("Expected node %d to contain %s, but was %s", i, result[i], node.Element())
		}
		i++
	}

}

func TestRevIterateAfterInsert(t *testing.T) {
	list := NewWithElements("John", "Ashling", "Finn")
	list.Insert("Daniella", 0)
	list.Insert("Anthony", 4)
	list.Insert("Peter", 2)
	// result is Daniella, John, Peter, Ashling, Finn, Anthony
	result := map[int]string{0: "Daniella", 1: "John", 2: "Peter", 3: "Ashling", 4: "Finn", 5: "Anthony"}
	var i int
	for node := list.Tail(); node != nil; node = node.Prev() {
		if result[5-i] != node.Element() {
			t.Fatalf("Expected node %d to contain %s, but was %s", i, result[i], node.Element())
		}
		i++
	}

}

func TestInsertFront(t *testing.T) {
	list := NewWithElements("John", "Ashling", "Finn")
	err := list.Insert("Anthony", 0)
	if err != nil {
		t.Fatal("Insertion failed")
	}

	if list.Head().Element() != "Anthony" {
		t.Fatal("Anthony was not inserted as the first node")
	}

}

func TestInsertFirst(t *testing.T) {

	list := NewWithElements("John", "Ashling", "Finn")

	err := list.Insert("Anthony", 1)
	if err != nil {
		t.Fatal("Insertion failed")
	}

	if list.Head().Next().Element() != "Anthony" {
		t.Fatal("Anthony was not inserted as the second node")
	}

}

func TestInsertLast(t *testing.T) {

	list := NewWithElements("John", "Ashling", "Finn")

	err := list.Insert("Anthony", 3)
	if err != nil {
		t.Fatal("Unable to add Anthony to the end of the list")
	}
	if list.Head().Next().Next().Next().Element() != "Anthony" {
		t.Fatal("Anthony was not added to the end of the list")
	}
	if list.Tail().Element() != "Anthony" {
		t.Fatal("The tail node was not updated properly")
	}
}

func TestInsertPastLast(t *testing.T) {
	list := NewWithElements("John", "Ashling", "Finn")

	err := list.Insert("Anthony", 100)
	if err == nil {
		t.Fatal("Should not be able to insert well past the end of the list")
	}

}

func TestDelete(t *testing.T) {
	list := NewWithElements("John", "Ashling", "Finn", "Anthony")
	err := list.Delete(0)
	if err != nil {
		t.Fatal("Failed to delete")
	}
	if list.Head().Element() != "Ashling" {
		t.Fatal("John was not deleted from the beginning of the list")
	}

	err = list.Delete(1)
	if err != nil {
		t.Fatal("Failed to delete")
	}
	if list.Head().Next().Element() != "Anthony" {
		t.Fatal("Finn was not deleted from the second position in the list")
	}

	err = list.Delete(1)
	if err != nil {
		t.Fatal("Failed to delete")
	}
	if list.Head().Element() != "Ashling" && list.Tail().Element() != "Ashling" {
		t.Fatal("Anthony was not deleted from the end of the list")
	}
	err = list.Delete(100)
	if err == nil {
		t.Fatal("Should not be able to delete from well past the end of the list")
	}

}

func TestSize(t *testing.T) {
	list := NewWithElements("John", "Ashling", "Finn")
	size := list.Size()
	if size != 3 {
		t.Fatal("Size of list should be 3")
	}
}
