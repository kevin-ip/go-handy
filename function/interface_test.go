package function

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Foo interface {
	Something() int
}

type foo struct {
	value int
}

func (f *foo) Something() int {
	return f.value
}

func TestIsNil(t *testing.T) {
	t.Run("some value to an interface should be true", func(t *testing.T) {
		var f Foo
		aFoo := &foo{value: 5}
		f = aFoo
		require.False(t, IsNil(f))
	})

	t.Run("nil value to an interface should be false", func(t *testing.T) {
		var f Foo
		require.True(t, IsNil(f))
	})

	t.Run("nil value to an interface should be false", func(t *testing.T) {
		var f Foo
		var aFoo *foo = nil
		f = aFoo
		require.True(t, IsNil(f))
	})
}
