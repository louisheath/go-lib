package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		strMap := map[string]struct{}{}
		strSet := New[string]()

		require.Equal(t, strMap, map[string]struct{}(strSet))
		require.Equal(t, Set[string](strMap), strSet)
	})

	t.Run("int", func(t *testing.T) {
		intMap := map[int]struct{}{}
		intSet := New[int]()

		require.Equal(t, intMap, map[int]struct{}(intSet))
		require.Equal(t, Set[int](intMap), intSet)
	})
}

func TestNewWithValues(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		strMap := map[string]struct{}{"a": {}, "B": {}}
		strSet := New("B", "a")

		require.Equal(t, strMap, map[string]struct{}(strSet))
		require.Equal(t, Set[string](strMap), strSet)
	})

	t.Run("int", func(t *testing.T) {
		intMap := map[int]struct{}{1: {}, 3: {}}
		intSet := New(3, 1)

		require.Equal(t, intMap, map[int]struct{}(intSet))
		require.Equal(t, Set[int](intMap), intSet)
	})
}

func TestSet(t *testing.T) {
	t.Parallel()

	strSet := New[string]()
	require.Len(t, strSet, 0)
	require.Equal(t, 0, strSet.Size())

	strSet.Add("z")
	require.Len(t, strSet, 1)
	require.Equal(t, 1, strSet.Size())
	require.True(t, strSet.Contains("z"))

	strSet.Add("x")
	require.Len(t, strSet, 2)
	require.Equal(t, 2, strSet.Size())

	strSet.Remove("z")
	require.Len(t, strSet, 1)
	require.Equal(t, 1, strSet.Size())
	require.False(t, strSet.Contains("z"))

	strSet.Remove("y")
	require.Len(t, strSet, 1)
	require.Equal(t, 1, strSet.Size())
	require.False(t, strSet.IsEmpty())

	strSet.Remove("x")
	require.Len(t, strSet, 0)
	require.Equal(t, 0, strSet.Size())
	require.True(t, strSet.IsEmpty())
}
