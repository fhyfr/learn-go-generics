package go_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Length :nodoc:
func Length[T interface{}]() {

}

// Width :nodoc:
func Width[W any](param W) W {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	string := Width("fireman")
	assert.Equal(t, string, "fireman")

	int := Width(29)
	assert.Equal(t, int, 29)

	var bool bool = Width[bool](false)
	assert.False(t, bool)
}
