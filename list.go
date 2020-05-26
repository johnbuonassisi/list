package list

import (
	"errors"
)

// Node represents an item in a linked list
type Node struct {
	prev    *Node  // points to the previous Node in the list
	next    *Node  // points to the next Node in the list
	element string // is the data carried by the Node
}

// NewNode creates a new node
func NewNode(prev *Node, next *Node, element string) Node {
	return Node{prev: prev, next: next, element: element}
}

// Prev returns a pointer to the previous Node
func (n *Node) Prev() *Node {
	return n.prev
}

// Next returns a pointer to the next Node
func (n *Node) Next() *Node {
	return n.next
}

// Element return the element of the Node, which is the main data carried by the Node
func (n *Node) Element() string {
	return n.element
}

// List implements a doubly linked list data structure
type List struct {
	head *Node
	tail *Node
}

// New creates an empty list
func New() List {
	return List{head: nil, tail: nil}
}

// NewWithElements creates a list containing the provided elements, in the order they are given
func NewWithElements(elements ...string) List {
	list := New()
	for _, element := range elements {
		list.Add(element)
	}
	return list
}

// Head returns a pointer to the first Node in the list
func (l *List) Head() *Node {
	return l.head
}

// Tail returns a pointer to the last Node in the list
func (l *List) Tail() *Node {
	return l.tail
}

// Add an element to the end of the list
func (l *List) Add(element string) {
	// create node thats previous node is the tail node and next node is nothing
	node := NewNode(l.tail, nil, element)
	// if head node is nil then set the new node to the head
	if l.head == nil {
		l.head = &node
	} else {
		l.tail.next = &node
	}
	// make the new node the tail node
	l.tail = &node
}

// Insert an element before an element at position pos
func (l *List) Insert(element string, pos int) error {

	// get the nodes before and after pos
	var prevNode *Node
	var replacementNode *Node
	var i int
	for node := l.head; node != nil; node = node.Next() {
		if i == pos-1 {
			prevNode = node
		}
		if i == pos {
			replacementNode = node
			break
		}
		i++
	}

	// error if attempting to insert past the end of the list
	if prevNode == nil && replacementNode == nil {
		return errors.New("Unable to find insertion position")
	}

	// when inserting after the first node in the list, set
	// the previous node to point to the new node, and set the
	// new node to point back to the previous node
	newNode := NewNode(nil, nil, element)
	if prevNode != nil {
		prevNode.next = &newNode
		newNode.prev = prevNode
	} else {
		// the newNode is being added to the front of the list, update the head pointer
		l.head = &newNode
	}

	// set the node after to point back to the new node, set the
	// next node of the new node to the node after the insertion position
	if replacementNode != nil {
		newNode.next = replacementNode
		replacementNode.prev = &newNode
	} else {
		// the newNode is being added to the end of the list, update the tail pointer
		l.tail = &newNode
	}

	return nil

}

// Delete an element at position pos
func (l *List) Delete(pos int) error {
	// Go to the node to be deleted
	var deleteNode *Node
	var i int
	for node := l.head; node != nil; node = node.Next() {
		if i == pos {
			deleteNode = node
			break
		}
		i++
	}

	if deleteNode == nil {
		return errors.New("Can't delete node, no node at specified position")
	}

	// Update prev and next pointers to pass over the deleted node
	if deleteNode.prev != nil {
		deleteNode.prev.next = deleteNode.next
	} else {
		l.head = deleteNode.next
	}

	// Set tail node if tail was deleted
	if deleteNode.next != nil {
		deleteNode.next.prev = deleteNode.prev
	} else {
		l.tail = deleteNode.prev
	}

	return nil
}

// Size of the list
func (l *List) Size() int {
	// we could calculate the size after each node is added to the list
	// but for simplicity lets just iterate the entire list each time this
	// function is called
	var size int
	for node := l.head; node != nil; node = node.Next() {
		size++
	}
	return size
}
