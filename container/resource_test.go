package container_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elemir/gloomo/container"
)

func TestResource(t *testing.T) {
	var res container.Resource[string]

	str, ok := res.Get()
	require.Equal(t, "", str)
	require.False(t, ok)

	res.Set("set")
	str, ok = res.Get()
	require.Equal(t, "set", str)
	require.True(t, ok)
}
