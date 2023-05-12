package collections_test

import (
	"testing"

	"github.com/barrettj12/collections"
	"github.com/stretchr/testify/assert"
)

func TestGetWrongTypePanics(t *testing.T) {
	wrapped := collections.Inject0[string, int]("string")
	assert.Equal(t, wrapped.Is0(), true)
	assert.Equal(t, wrapped.Is1(), false)
	assert.Panics(t, func() { wrapped.Get1() })
}
