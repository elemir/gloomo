package container_test

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elemir/gloomo/container"
	"github.com/elemir/gloomo/id"
)

func TestSparseArray(t *testing.T) {
	var arr container.SparseArray[string]

	t.Run("insert", func(t *testing.T) {
		expected := map[id.ID]string{
			10: "meow",
			1:  "hello",
			5:  "elephant",
		}

		arr.Set(10, "meow")
		arr.Set(1, "hello")
		arr.Set(5, "elephant")

		actual := maps.Collect(arr.Items())

		require.Equal(t, expected, actual)

		value, ok := arr.Get(10)
		require.Equal(t, "meow", value)
		require.True(t, ok)

		value, ok = arr.Get(11)
		require.Equal(t, "", value)
		require.False(t, ok)
	})

	t.Run("update", func(t *testing.T) {
		expected := map[id.ID]string{
			10: "meow",
			1:  "goodbuy",
			5:  "elephant",
		}

		arr.Set(1, "goodbuy")

		actual := maps.Collect(arr.Items())

		require.Equal(t, expected, actual)
	})

	t.Run("delete", func(t *testing.T) {
		expected := map[id.ID]string{
			1: "goodbuy",
			5: "elephant",
		}

		arr.Delete(10)

		actual := maps.Collect(arr.Items())

		require.Equal(t, expected, actual)
	})

	t.Run("delete-not-exists", func(t *testing.T) {
		expected := map[id.ID]string{
			1: "goodbuy",
			5: "elephant",
		}

		arr.Delete(14)

		actual := maps.Collect(arr.Items())

		require.Equal(t, expected, actual)
	})

	t.Run("break", func(t *testing.T) {
		var i int

		for range arr.Items() {
			i++

			break
		}

		require.Equal(t, 1, i)
	})

	t.Run("delete-all", func(t *testing.T) {
		arr.Delete(1)
		arr.Delete(5)

		require.Empty(t, maps.Collect(arr.Items()))
	})

	t.Run("work-with-ints", func(t *testing.T) {
		var intArr container.SparseArray[int]

		expected := map[id.ID]int{
			10: 20,
			1:  11,
			5:  15,
		}

		intArr.Set(10, 20)
		intArr.Set(1, 11)
		intArr.Set(5, 15)

		actual := maps.Collect(intArr.Items())
		require.Equal(t, expected, actual)

		value, ok := intArr.Get(10)
		require.Equal(t, 20, value)
		require.True(t, ok)
	})
}
