package list

import (
	"fmt"
)

// Node contains value to store the data
// prev point to the previous node in linked list
// next point to the next node in linked list
type node struct {
	value interface{}
	prev  *node
	next  *node
}

// DoubleLinkedList contains head pointer points to the first node of linked list
// tail pointer points to the last node of linked list
// size store the current number item in linked list
type DoubleLinkedList struct {
	head *node
	tail *node
	size int
}

// New returns a pointer to new DoubleLinkedList object
func New() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

// NewWithItems returns a pointer to new DoubleLinkedList object
func NewWithItems(vv ...interface{}) *DoubleLinkedList {
	return &DoubleLinkedList{}
}

// Size returns the current size of linked list
func (l *DoubleLinkedList) Size() int {
	return l.size
}

// IsEmpty returns true if list is empty
// returns false if list not empty
func (l *DoubleLinkedList) IsEmpty() bool {
	return l.size == 0
}

// String returns string form of linked list
func (l *DoubleLinkedList) String() string {
	if l.IsEmpty() {
		return "[]"
	}
	var str string = "["
	trav := l.head
	for trav != nil {
		str += fmt.Sprintf("%v", trav.value)
		trav = trav.next
		if trav != nil {
			str += ", "
		} else {
			str += "]"
		}
	}

	return str
}

// Add append a list of item to linked list
func (l *DoubleLinkedList) Add(v interface{}) {
	l.AddLast(v)
}

// AddLast append a item to linked list
func (l *DoubleLinkedList) AddLast(v interface{}) {
	newNode := &node{value: v, next: nil, prev: l.tail}
	if l.IsEmpty() {
		l.head, l.tail = newNode, newNode
	} else {
		l.tail.next, l.tail = newNode, newNode
	}

	l.size++
}

// AddFirst adds a node to the beginning of linked list
func (l *DoubleLinkedList) AddFirst(v interface{}) {
	newNode := &node{value: v, next: l.head, prev: nil}
	if l.IsEmpty() {
		l.head, l.tail = newNode, newNode
	} else {
		l.head.prev, l.head = newNode, newNode
	}

	l.size++
}

// AddAt adds value to a specific index
func (l *DoubleLinkedList) AddAt(index int, v interface{}) error {
	if index > l.size || index < 0 {
		return fmt.Errorf("Illegal Index %v", index)
	}

	if index == 0 {
		l.AddFirst(v)
		return nil
	}

	if index == l.size {
		l.AddLast(v)
		return nil
	}

	trav := l.head
	for i := 0; i < index-1; i++ {
		trav = trav.next
	}

	newNode := &node{value: v, next: trav.next, prev: trav}
	trav.next = newNode
	l.size++

	return nil
}

// PeekFirst returns value of first item
// returns List Empty error if list is empty
func (l *DoubleLinkedList) PeekFirst() (interface{}, error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("List Empty")
	}

	return l.head.value, nil
}

// PeekLast returns value of last item
// returns List Empty error if list is empty
func (l *DoubleLinkedList) PeekLast() (interface{}, error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("List Empty")
	}

	return l.tail.value, nil
}

// RemoveFirst removes and returns first item value
// returns List Empty error if list is empty
func (l *DoubleLinkedList) RemoveFirst() (interface{}, error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("List Empty")
	}

	v := l.head.value
	l.head = l.head.next
	l.size--

	if l.IsEmpty() {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	return v, nil
}

// RemoveLast removes and returns last item value
// returns List Empty error if list is empty
func (l *DoubleLinkedList) RemoveLast() (interface{}, error) {
	if l.IsEmpty() {
		return nil, fmt.Errorf("List Empty")
	}

	v := l.tail.value
	l.tail = l.tail.prev
	l.size--

	if l.IsEmpty() {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	return v, nil
}

// IndexOf returns the index of value
// returns -1 if dont exist or list is empty
func (l *DoubleLinkedList) IndexOf(v interface{}) int {
	_, index := l.find(v)
	return index
}

// Remove removes the specific value
// returns List Empty error if list is empty
func (l *DoubleLinkedList) Remove(v interface{}) error {
	if l.IsEmpty() {
		return fmt.Errorf("List Empty")
	}

	n, _ := l.find(v)
	_, err := l.remove(n)

	return err
}

// RemoveAt removes value at the specific index
// returns List Empty error if list is empty
func (l *DoubleLinkedList) RemoveAt(index int) (interface{}, error) {
	if index >= l.size || index < 0 {
		return nil, fmt.Errorf("Illegal Index %v", index)
	}

	var trav *node
	if index < l.size/2 {
		trav = l.head
		for i := 0; i != index; i++ {
			trav = trav.next
		}
	} else {
		trav = l.tail
		for i := l.size - 1; i != index; i-- {
			trav = trav.prev
		}
	}

	return l.remove(trav)
}

// Contains returns true if value is existed
// returns false other cases
func (l *DoubleLinkedList) Contains(v interface{}) bool {
	return l.IndexOf(v) != -1
}

// Clear empties linked list
func (l *DoubleLinkedList) Clear() {
	trav := l.head
	for trav != nil {
		next := trav.next
		trav.prev, trav.next, trav.value = nil, nil, nil
		trav = next
	}
	l.head, l.tail, l.size = nil, nil, 0
}

func (l *DoubleLinkedList) remove(n *node) (interface{}, error) {
	if n.prev == nil {
		return l.RemoveFirst()
	}

	if n.next == nil {
		return l.RemoveLast()
	}

	v := n.value
	n.prev.next, n.next.prev = n.next, n.prev
	l.size--

	n.prev, n.next, n = nil, nil, nil

	return v, nil
}

func (l *DoubleLinkedList) find(v interface{}) (*node, int) {
	trav, c := l.head, 0
	for trav != nil {
		if trav.value == v {
			return trav, c
		}
		trav = trav.next
		c++
	}

	return nil, -1
}
