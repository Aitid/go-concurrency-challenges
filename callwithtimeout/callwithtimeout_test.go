package callwithtimeout

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCallWithTimeout(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		ctx := context.Background()
		timeout := time.Duration(1 * time.Second)
		fn := func() (string, error) {
			return "data from API", nil
		}
		result, _ := CallWithTimeout(ctx, fn, timeout)

		assert.Equal(t, "data from API", result)
	})
	t.Run("timeout", func(t *testing.T) {
		ctx := context.Background()
		timeout := 10 * time.Millisecond

		fn := func() (string, error) {
			time.Sleep(100 * time.Millisecond)
			return "data", nil
		}

		result, err := CallWithTimeout(ctx, fn, timeout)

		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
	t.Run("context canceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		fn := func() (string, error) {
			time.Sleep(50 * time.Millisecond)
			return "data", nil
		}

		result, err := CallWithTimeout(ctx, fn, time.Second)

		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.Canceled)
	})
	t.Run("zero timeout", func(t *testing.T) {
		ctx := context.Background()

		fn := func() (string, error) {
			return "data", nil
		}

		result, err := CallWithTimeout(ctx, fn, 0)

		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
	t.Run("function returns error", func(t *testing.T) {
		ctx := context.Background()

		expectedErr := errors.New("api error")
		fn := func() (string, error) {
			return "", expectedErr
		}

		result, err := CallWithTimeout(ctx, fn, time.Second)

		assert.Empty(t, result)
		assert.Equal(t, expectedErr, err)
	})
	t.Run("context canceled during execution", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		fn := func() (string, error) {
			time.Sleep(100 * time.Millisecond)
			return "data", nil
		}

		go func() {
			time.Sleep(10 * time.Millisecond)
			cancel()
		}()

		result, err := CallWithTimeout(ctx, fn, time.Second)

		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.Canceled)
	})
	t.Run("completion near deadline", func(t *testing.T) {
		ctx := context.Background()
		timeout := 50 * time.Millisecond

		fn := func() (string, error) {
			time.Sleep(45 * time.Millisecond)
			return "data", nil
		}

		result, err := CallWithTimeout(ctx, fn, timeout)

		assert.NoError(t, err)
		assert.Equal(t, "data", result)
	})
	t.Run("parent context already canceled", func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		cancel()

		fn := func() (string, error) {
			return "data", nil
		}

		result, err := CallWithTimeout(parent, fn, time.Second)

		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.Canceled)
	})
}
