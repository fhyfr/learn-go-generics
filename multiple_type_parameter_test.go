package go_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// MultipleTypeParam :nodoc:
func MultipleTypeParam[MTP1 any, MTP2 interface{}](param1 MTP1, param2 MTP2) (MTP1, MTP2) {

	return param1, param2
}

func TestMultipleTypeParam(t *testing.T) {
	name, age := MultipleTypeParam[string, int]("Firman", 23)
	assert.Equal(t, name, "Firman")
	assert.Equal(t, age, 23)
}
