package caution

// This file isn't guarded with either the "private" or "!private" build tags, it will be
// compiled in either event.

// Thinger should have two implementations, one for use with the private Go module, and
// another placeholder used when the private Go module isn't available.
type Thinger interface {
	Thing() string
}
