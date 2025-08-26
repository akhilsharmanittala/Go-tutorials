// linked list is a collection of elements called nodes
// each node contains a value and a pointer to the next node in the sequence
// the first node is called the head and the last node is called the tail
// the tail node points to nil, indicating the end of the list

// head                                           //tail
// [A| next] ---> ][B| next] ---> ][C| next] ---> [D| nil]
package main

import "fmt"

// Node represents a single node in the list
type Node struct {
	value int
	next  *Node
}

// List represents the singly linked list
type List struct {
	head *Node
	tail *Node
}

// Add adds a new node at the end
func (l *List) Add(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
}

// AddAtBeginning adds a node at the start of the list
func (l *List) AddAtBeginning(value int) {
	newNode := &Node{value: value, next: l.head}
	if l.head == nil {
		l.tail = newNode
	}
	l.head = newNode
}

// AddAtEnd adds a node at the end of the list
func (l *List) AddAtEnd(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
}

// AddAtPosition inserts at a specific position (0-based index)
func (l *List) AddAtPosition(value, position int) {
	newNode := &Node{value: value}
	if position == 0 {
		l.AddAtBeginning(value)
		return
	}
	current := l.head
	for i := 0; current != nil && i < position-1; i++ {
		current = current.next
	}
	if current == nil {
		fmt.Println("Position out of range")
		return
	}
	newNode.next = current.next
	current.next = newNode
	if newNode.next == nil {
		l.tail = newNode
	}
}

// DeleteAtBeginning removes the first node
func (l *List) DeleteAtBeginning() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
}

// DeleteAtEnd removes the last node
func (l *List) DeleteAtEnd() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.head.next == nil {
		l.head, l.tail = nil, nil
		return
	}
	current := l.head
	for current.next != l.tail {
		current = current.next
	}
	current.next = nil
	l.tail = current
}

// DeleteByValue deletes the first node with the given value
func (l *List) DeleteByValue(value int) {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.head.value == value {
		l.DeleteAtBeginning()
		return
	}
	current := l.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next == nil {
		fmt.Println("Value not found")
		return
	}
	if current.next == l.tail {
		l.tail = current
	}
	current.next = current.next.next
}

// Search checks if a value exists in the list
func (l *List) Search(value int) bool {
	current := l.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

// Length returns the number of nodes in the list
func (l *List) Length() int {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Reverse reverses the list
func (l *List) Reverse() {
	var prev *Node
	current := l.head
	l.tail = l.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

// Traverse prints the linked list
func (l *List) Traverse() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func SingleLinkedList() {
	list := &List{}

	fmt.Println("Adding nodes at the end:")
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Traverse()

	fmt.Println("\nAdding at beginning:")
	list.AddAtBeginning(5)
	list.Traverse()

	fmt.Println("\nAdding at end:")
	list.AddAtEnd(40)
	list.Traverse()

	fmt.Println("\nAdding at position 2:")
	list.AddAtPosition(15, 2)
	list.Traverse()

	fmt.Println("\nSearching for value 20:", list.Search(20))
	fmt.Println("Searching for value 99:", list.Search(99))

	fmt.Println("\nDeleting first node:")
	list.DeleteAtBeginning()
	list.Traverse()

	fmt.Println("\nDeleting last node:")
	list.DeleteAtEnd()
	list.Traverse()

	fmt.Println("\nDeleting value 15:")
	list.DeleteByValue(15)
	list.Traverse()

	fmt.Println("\nReversing list:")
	list.Reverse()
	list.Traverse()

	fmt.Println("\nLength of list:", list.Length())
}
