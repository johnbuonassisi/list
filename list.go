package list

// Node represents an item in a linked list
type Node struct {
	prev *Node       // points to the previous Node in the list
	next *Node       // points to the next Node in the list
	Data interface{} // is the data carried by the Node
	list *List       // the list the node belongs to
}

// Prev returns a pointer to the previous Node
func (n *Node) Prev() *Node {
	if n.list != nil && n.prev != &n.list.root {
		return n.prev
	}
	return nil
}

// Next returns a pointer to the next Node/
func (n *Node) Next() *Node {
	// return nil when root element is hit
	if n.list != nil && n.next != &n.list.root {
		return n.next
	}
	return nil
}

// List implements a doubly linked list data structure
type List struct {
	root Node
	Size int
}

// New creates an empty list
func New() *List {
	// new returns a pointer to a List with all variables set to 0 value
	return new(List).Init()
}

// Init the root node to point to itself, ensures size is 0
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.Size = 0
	return l
}

// lazyInit will initialize a list that has not already been initialized
// this is used if a programmer allocates a list like
// var l = new(List) or l := List{} and then immediately tries to use it
// like l.Add("test"). lazyInit is called inside Add, which will call
// Init() only if l.root.next == nil, which is only true if Init is never
// called.
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// Head returns a pointer to the first Node in the list, nil if empty
func (l *List) Head() *Node {
	// when the size is 0 return nil because the next node of the root
	// node is itself
	if l.Size == 0 {
		return nil
	}
	return l.root.next
}

// Tail returns a pointer to the last Node in the list, nil if empty
func (l *List) Tail() *Node {
	// when the size is 0 return nil because the next node of the root
	// code is itself
	if l.Size == 0 {
		return nil
	}
	return l.root.prev
}

// AddToBack a new node with data d into the front of the list
func (l *List) AddToBack(d interface{}) *Node {
	// lazyInit just in case
	// this will guarantee that the next and prev nodes of root are not nil
	l.lazyInit()

	// insert the node to the end of the list
	return l.insertData(d, l.root.prev)
}

// AddToFront inserts a new node with data d into the front of the list
func (l *List) AddToFront(d interface{}) *Node {
	l.lazyInit()
	return l.insertData(d, &l.root)
}

// AddBefore inserts a new node with data d before at
func (l *List) AddBefore(d interface{}, at *Node) *Node {
	l.lazyInit()
	return l.insertData(d, at.prev)
}

// AddAfter inserts a new node with data d after at
func (l *List) AddAfter(d interface{}, at *Node) *Node {
	l.lazyInit()
	return l.insertData(d, at)
}

// insertData is a helper function to wrap d in a node and insert
// it into a list after at
func (l *List) insertData(d interface{}, at *Node) *Node {
	l.lazyInit()
	return l.insert(&Node{Data: d}, at)
}

// insert node n after at, increments l.size, returns n
// Generic function that should be used by any exported method that
// add an node to the list
func (l *List) insert(n, at *Node) *Node {
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.Size++
	return n
}

// Delete node n from the list and return its value
func (l *List) Delete(n *Node) interface{} {
	if n.list == l {
		l.remove(n)
	}
	return n.Data
}

// remove n from the list
func (l *List) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	n.list = nil
	l.Size--
}
