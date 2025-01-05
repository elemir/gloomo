package container_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elemir/gloomo/container"
)

func TestZero(t *testing.T) {
	t.Parallel()

	require.Equal(t, "", container.Zero[string](), "")
	require.Equal(t, 0, container.Zero[int]())
	//nolint: testifylint
	require.Equal(t, 0.0, container.Zero[float64]())
	require.Equal(t, struct{}{}, container.Zero[struct{}]())
	require.Equal(t, (*int)(nil), container.Zero[*int]())
	require.Nil(t, container.Zero[any]())
}
