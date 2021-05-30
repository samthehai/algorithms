package compstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	compStr := NewCompressString()

	assert.Equal(t, "", compStr.Compress(""))
	assert.Equal(t, "AABBCC", compStr.Compress("AABBCC"))
	assert.Equal(t, "A3BC2D4E", compStr.Compress("AAABCCDDDDE"))
	assert.Equal(t, "BA3C2D4", compStr.Compress("BAAACCDDDD"))
	assert.Equal(t, "A3BA2C2D4", compStr.Compress("AAABAACCDDDD"))
}
