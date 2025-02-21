package collectiontest

import (
	"testing"

	"github.com/kevin-ip/go-handy/collection"
)

func TestLinkedDeque(t *testing.T) {
	RunDequeTests(t, func() collection.Deque[int] {
		return collection.NewLinkedDeque[int]()
	})
}
