package context

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDetached(t *testing.T) {
	t.Parallel()

	t.Run("exposes values", func(t *testing.T) {
		parent := context.Background()
		parent = context.WithValue(parent, "first", "a")

		detached := NewDetached(parent)
		detached = context.WithValue(detached, "second", "b")

		require.Equal(t, "a", detached.Value("first"))
		require.Equal(t, "b", detached.Value("second"))
	})

	t.Run("child ignores ignores parent deadline", func(t *testing.T) {
		parent := context.Background()
		_, deadlineExists := parent.Deadline()
		require.False(t, deadlineExists)

		parent, cancel := context.WithTimeout(parent, 50*time.Millisecond)
		defer cancel()
		_, deadlineExists = parent.Deadline()
		require.True(t, deadlineExists)

		detached := NewDetached(parent)
		_, deadlineExists = detached.Deadline()
		require.False(t, deadlineExists)

		require.Eventually(t, func() bool {
			return errors.Is(parent.Err(), context.DeadlineExceeded)
		}, 300*time.Millisecond, 10*time.Millisecond)
		require.NoError(t, detached.Err())
	})

	t.Run("child ignores parent cancel", func(t *testing.T) {
		parent := context.Background()

		parent, cancel := context.WithCancel(parent)
		detached := NewDetached(parent)

		cancel()
		require.ErrorIs(t, parent.Err(), context.Canceled)
		require.NoError(t, detached.Err())
	})

	t.Run("parent ignores ignores child deadline", func(t *testing.T) {
		parent := context.Background()
		_, deadlineExists := parent.Deadline()
		require.False(t, deadlineExists)

		detached := NewDetached(parent)
		_, deadlineExists = detached.Deadline()
		require.False(t, deadlineExists)

		detached, cancel := context.WithTimeout(detached, 50*time.Millisecond)
		defer cancel()
		_, deadlineExists = detached.Deadline()
		require.True(t, deadlineExists)

		require.Eventually(t, func() bool {
			return errors.Is(detached.Err(), context.DeadlineExceeded)
		}, 300*time.Millisecond, 10*time.Millisecond)
		require.NoError(t, parent.Err())
	})

	t.Run("parent ignores child cancel", func(t *testing.T) {
		parent := context.Background()

		detached := NewDetached(parent)
		detached, cancel := context.WithCancel(detached)

		cancel()
		require.ErrorIs(t, detached.Err(), context.Canceled)
		require.NoError(t, parent.Err())
	})
}
