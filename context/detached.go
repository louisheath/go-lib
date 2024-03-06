package context

import (
	"context"
	"time"
)

type detachedContext struct {
	context.Context
	parent context.Context
}

// NewDetached returns a new context with the same values as the parent. This can be used
// to clone a context without inheriting the cancellation properities of the parent,
// meaning if the parent is cancelled, the child is not. Likewise, the child context does
// not cancel the parent. This can be useful when performing operations asyncronously.
//
// Note that this goes against the designed behaviour of Go contexts and therefore this
// function should be used with caution. See similar:
// https://github.com/golang/tools/blob/master/internal/xcontext/xcontext.go
func NewDetached(parent context.Context) context.Context {
	return &detachedContext{
		Context: context.Background(),
		parent:  parent,
	}
}

func (d detachedContext) Value(key any) any {
	return d.parent.Value(key)
}

// NewDetachedWithTimeout returns a new detached child context that cancels after the
// provided duration. It is not cancelled by the parent context.
func NewDetachedWithTimeout(
	parent context.Context, dur time.Duration,
) (context.Context, context.CancelFunc) {
	detached := NewDetached(parent)
	return context.WithTimeout(detached, dur)
}
