package container_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elemir/gloomo/container"
)

func TestSet(t *testing.T) {
	var set container.Set

	t.Run("delete-from-empty", func(_ *testing.T) {
		set.Delete(0)
	})

	t.Run("get-from-empty", func(t *testing.T) {
		require.False(t, set.Get(0))
	})

	t.Run("add", func(t *testing.T) {
		set.Add(0)
		require.True(t, set.Get(0))
		require.False(t, set.Get(1))
	})

	t.Run("delete", func(t *testing.T) {
		set.Delete(0)
		require.False(t, set.Get(0))
	})
}
