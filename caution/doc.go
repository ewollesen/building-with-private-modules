// package caution will optionally link to private Go modules.
//
// Building with a tag of "private" will produce a binary that links to a Tidepool partner's
// private SDK. Building without the "private" tag will instead build placeholders. See the
// [Go build constraints documentation] for details.
//
// # !!! ALL REFERENCES TO PRIVATE GO MODULES MUST BE MADE FROM GUARDED FILES !!!
//
// Doing so ensures that the default build—one with no tags specified—will *NOT* link to
// private Go packages.
//
// Any symbols exported by this package, should have two versions: one that is guarded with
// the "private" tag and references the private Go module, and another that is guarded with
// the "!private" tag and provides placeholder implementations.
//
// # Conventions
//
// A suggested convention to help prevent accidental code leakage follows:
//
//	foo/
//	   + foo_common.go  -- Contains common interface definitions. Unguarded.
//	   + foo_private.go -- Contains real implementations. Guarded with "//go:build private".
//	   + foo_safe.go    -- Contains placeholder implementations. Guarded with "//go:build !private".
//
// # Future Considerations
//
// For the future, it might be possible to use something like the [impl] package to generate
// public placeholder implementations of private integration interfaces. Stubbing out
// placholder struct types could be annoying. When we have a more concrete use case, we can
// evaluate the feasibility.
//
// Another future possibility, would be to encrypt private.go via [git-crypt], such that
// even the hint of the integration is ket private. Hopefully that won't be necessary.
//
// If the convention above is implemented and found useful, [Git hooks] could be written to
// check for the existence of tags in files that match the convention.
//
// [impl]: https://pkg.go.dev/github.com/josharian/impl
// [git-crypt]: https://www.agwa.name/projects/git-crypt/
// [Go build constraints documentation]: https://pkg.go.dev/cmd/go#hdr-Build_constraints
// [Git hooks]: https://git-scm.com/book/ms/v2/Customizing-Git-Git-Hooks
package caution
