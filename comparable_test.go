package go_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// IsSame :nodoc:
func IsSame[T comparable](param1, param2 T) bool {
	return param1 == param2
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Firman", "Firman"))
	assert.False(t, IsSame[string]("firman", "Firman"))
	assert.Equal(t, IsSame[int](1, 2), false)
}
