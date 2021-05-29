package minheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	var nilInt *int

	heap := NewMinHeap()
	assert.Equal(t, heap.PeekMin(), nilInt)
	assert.Equal(t, heap.ExtractMin(), nilInt)

	heap.Insert(20)
	assert.Equal(t, 20, heap.array[0])

	heap.Insert(5)
	assert.Equal(t, 5, heap.array[0])
	assert.Equal(t, 20, heap.array[1])

	heap.Insert(15)
	assert.Equal(t, 5, heap.array[0])
	assert.Equal(t, 20, heap.array[1])
	assert.Equal(t, 15, heap.array[2])

	heap.Insert(22)
	assert.Equal(t, 5, heap.array[0])
	assert.Equal(t, 20, heap.array[1])
	assert.Equal(t, 15, heap.array[2])
	assert.Equal(t, 22, heap.array[3])

	heap.Insert(40)
	assert.Equal(t, 5, heap.array[0])
	assert.Equal(t, 20, heap.array[1])
	assert.Equal(t, 15, heap.array[2])
	assert.Equal(t, 22, heap.array[3])
	assert.Equal(t, 40, heap.array[4])

	heap.Insert(3)
	assert.Equal(t, 3, heap.array[0])
	assert.Equal(t, 20, heap.array[1])
	assert.Equal(t, 5, heap.array[2])
	assert.Equal(t, 22, heap.array[3])
	assert.Equal(t, 40, heap.array[4])
	assert.Equal(t, 15, heap.array[5])
}
