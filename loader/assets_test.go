package loader_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elemir/gloomo/loader"
)

func TestAssets(t *testing.T) {
	var assets loader.Assets[string]

	t.Run("add-tasks", func(t *testing.T) {
		expected := []string{
			"path1",
			"path2",
		}

		asset, ok := assets.Load("path1")
		require.Empty(t, asset)
		require.False(t, ok)

		assets.Load("path2")
		require.Empty(t, asset)
		require.False(t, ok)

		actual := slices.Collect(assets.NotLoadedPaths())

		require.Equal(t, expected, actual)
	})

	t.Run("load-twice", func(t *testing.T) {
		asset, ok := assets.Load("path2")
		require.Empty(t, asset)
		require.False(t, ok)
	})

	t.Run("put-asset", func(t *testing.T) {
		expected := []string{
			"path2",
		}

		assets.Put("path1", "result")

		asset, ok := assets.Load("path1")
		require.Equal(t, "result", asset)
		require.True(t, ok)

		actual := slices.Collect(assets.NotLoadedPaths())

		require.Equal(t, expected, actual)
	})

	t.Run("replace-asset", func(t *testing.T) {
		assets.Put("path1", "replace")

		asset, ok := assets.Load("path1")
		require.Equal(t, "replace", asset)
		require.True(t, ok)
	})

	t.Run("break", func(t *testing.T) {
		var i int

		for range assets.NotLoadedPaths() {
			i++

			break
		}

		require.Equal(t, 1, i)
	})

	t.Run("finish-all", func(t *testing.T) {
		assets.Put("path2", "asset2")

		require.Empty(t, slices.Collect(assets.NotLoadedPaths()))
	})
}
