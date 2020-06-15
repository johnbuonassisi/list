A doubly linked list, implemented in Go with insertion, removal, and traversal operations.

Each record of a linked list is called a node

Each node contains a pointer to the next and revious node, and the data of the node itself

The list always contains an empty root node whose next node points to the first node and whose
previous node points to the last node.

Implementation closely follows container/list from the Go Standard Library
