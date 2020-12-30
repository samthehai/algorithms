package dynamicarray

import (
	"fmt"
)

// DynamicArray contains one sile to save  the items
// length is logical length, the length that user think it is
// capacity is the actual array size
// the reason why need length and capacity field is because below:
// https://stackoverflow.com/questions/26634554/go-multiple-len-calls-vs-performance
type DynamicArray struct {
	arr      []interface{}
	size     int
	capacity int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New instantiates a new dynamic array default will have capacity 16
func New() (*DynamicArray, error) {
	return NewWithCapacity(16)
}

// NewWithCapacity returns an Dynamic Array with specific capacity
// returns error when capacity parameter is illegal (below 0)
func NewWithCapacity(capacity int) (*DynamicArray, error) {
	if capacity < 0 {
		return nil, fmt.Errorf("Illegal Capacity: %v", capacity)
	}

	dynamicArray := &DynamicArray{
		capacity: capacity,
		arr:      make([]interface{}, 0, capacity),
	}

	return dynamicArray, nil
}

// Size returns the current length of Dynamic Array
func (a *DynamicArray) Size() int {
	return a.size
}

// IsEmpty returns true if array is empty
// returns false if array is not empty
func (a *DynamicArray) IsEmpty() bool {
	return a.Size() == 0
}

// Get returns the item of array at index
// returns error when index is out of bounds
func (a *DynamicArray) Get(index int) (interface{}, error) {
	if index >= a.size || index < 0 {
		return nil, fmt.Errorf("Index %v Is Out Of Bounds", index)
	}

	return a.arr[index], nil
}

// Set the item at index to value
// returns error when index is out of bounds
func (a *DynamicArray) Set(index int, v interface{}) error {
	if index >= a.size || index < 0 {
		return fmt.Errorf("Index %v Is Out Of Bounds", index)
	}

	a.arr[index] = v

	return nil
}

// Clear remove all the item in array
// reset length to zero, keep capacity as current
func (a *DynamicArray) Clear() {
	for i := 0; i < a.size; i++ {
		a.arr[i] = nil
	}

	a.size = 0
}

// Add item to array, double size if it reach limited
func (a *DynamicArray) Add(v interface{}) {
	a.grow()
	a.arr = append(a.arr, v)
	a.size = a.size + 1
}

// RemoveAt item from array, shrink size if necessary
func (a *DynamicArray) RemoveAt(index int) (interface{}, error) {
	if index >= a.size || index < 0 {
		return nil, fmt.Errorf("Index %v Is Out Of Bounds", index)
	}

	removeItem := a.arr[index]

	copy(a.arr[index:], a.arr[index+1:])
	a.arr[a.size-1] = nil
	a.arr = a.arr[:a.size-1]
	a.size--

	a.shrink()

	return removeItem, nil
}

// Remove item from array
func (a *DynamicArray) Remove(v interface{}) bool {
	index := a.IndexOf(v)
	if index == -1 {
		return false
	}

	if _, err := a.RemoveAt(index); err != nil {
		return false
	}

	return true
}

// IndexOf returns index of item in array
// returns -1 if not found
func (a *DynamicArray) IndexOf(v interface{}) int {
	for i := 0; i < a.size; i++ {
		if v == a.arr[i] {
			return i
		}
	}
	return -1
}

// Contains checks if element is present in the array.
func (a *DynamicArray) Contains(v interface{}) bool {
	return a.IndexOf(v) != -1
}

// Contains checks if element is present in the array.
func (a *DynamicArray) String() string {
	if a.size == 0 {
		return "[]"
	}
	var str string
	for i := 0; i < a.size; i++ {
		str += fmt.Sprintf("[%v]", a.arr[i])
		if i < a.size-1 {
			str += ","
		}
	}

	return str
}

// expand method double capacity of Dynamic Array
// it's current size. Fills the blanks with nil values
func (a *DynamicArray) grow() {
	if a.size+1 >= a.capacity {
		a.updateCapacity(int(float32(a.capacity) * growthFactor))
	}
}

// shrink the array when size is at shrinkFactor * capacity
func (a *DynamicArray) shrink() {
	if shrinkFactor == 0.0 {
		return
	}

	if a.size <= int(float32(a.capacity)*shrinkFactor) {
		a.updateCapacity(a.size)
	}
}

func (a *DynamicArray) updateCapacity(capacity int) {
	newArr := make([]interface{}, a.size, capacity)
	copy(newArr, a.arr)
	a.arr = newArr
	a.capacity = capacity
}
