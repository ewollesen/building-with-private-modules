//go:build !private

package caution

// In general, I don't expect us to write tests for the placeholders, but in case we did...

import "testing"

func TestSafeThinger(t *testing.T) {
	thinger := NewThinger()
	if got := thinger.Thing(); got != "placeholder thinger" {
		t.Errorf("expected placeholder Thinger, got %q", got)
	}
}
