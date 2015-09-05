package emoji

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestContains ...
func TestContains(t *testing.T) {
	assert := assert.New(t)

	msg1 := "Hello!"
	msg2 := "Hello! 😁"
	assert.False(Contains(msg1))
	assert.True(Contains(msg2))
}
