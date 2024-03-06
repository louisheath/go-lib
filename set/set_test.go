package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		strSet := New[string]()
		require.NotNil(t, strSet)
		require.Len(t, strSet, 0)
	})

	t.Run("int", func(t *testing.T) {
		intSet := New[int]()
		require.NotNil(t, intSet)
		require.Len(t, intSet, 0)
	})
}

func TestNewWithValues(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		strSet := New("B", "a")
		require.NotNil(t, strSet)
		require.Len(t, strSet, 2)
	})

	t.Run("int", func(t *testing.T) {
		intSet := New(3, 1)
		require.NotNil(t, intSet)
		require.Len(t, intSet, 2)
	})
}

func TestSet(t *testing.T) {
	t.Parallel()

	strSet := New[string]()
	require.Equal(t, strSet.Size(), 0)
	require.Equal(t, 0, strSet.Size())

	strSet.Add("z")
	require.Equal(t, 1, strSet.Size())
	require.True(t, strSet.Contains("z"))

	strSet.Add("x")
	require.Equal(t, 2, strSet.Size())

	strSet.Remove("z")
	require.Equal(t, 1, strSet.Size())
	require.False(t, strSet.Contains("z"))

	strSet.Remove("y")
	require.Equal(t, 1, strSet.Size())
	require.False(t, strSet.IsEmpty())

	strSet.Remove("x")
	require.Equal(t, 0, strSet.Size())
	require.True(t, strSet.IsEmpty())
}

func TestNil(t *testing.T) {
	t.Parallel()

	t.Run("ToSlice does not panic", func(t *testing.T) {
		var s Set[string]
		require.Empty(t, s.ToSlice())
	})

	t.Run("Add panics", func(t *testing.T) {
		var s Set[string]
		require.Panics(t, func() {
			s.Add("a")
		})
	})

	t.Run("AddAll panic", func(t *testing.T) {
		var s Set[string]
		require.Panics(t, func() {
			s.AddAll("a", "b")
		})
	})

	t.Run("Remove does not panic", func(t *testing.T) {
		var s Set[string]
		s.Remove("a")
		require.Equal(t, 0, s.Size())
	})

	t.Run("Size does not panic", func(t *testing.T) {
		var s Set[string]
		size := s.Size()
		require.Equal(t, 0, size)
	})

	t.Run("Contains does not panic", func(t *testing.T) {
		var s Set[string]
		contains := s.Contains("a")
		require.False(t, contains)
	})

	t.Run("IsEmpty does not panic", func(t *testing.T) {
		var s Set[string]
		isEmpty := s.IsEmpty()
		require.True(t, isEmpty)
	})
}
