//go:build private

package caution

// There would likely be tests here, that would only run if the build tag were provided.

import "testing"

func TestPrivateThinger(t *testing.T) {
	thinger := NewThinger()
	if got := thinger.Thing(); got != "This calls the real partner thinger." {
		t.Errorf("expected real Thinger, got %q", got)
	}
}
