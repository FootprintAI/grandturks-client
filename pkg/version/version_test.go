package version

import (
	"strings"
	"testing"

	goversion "github.com/hashicorp/go-version"
)

// The tests in this file exist because of grandturks-client#22: consumers pin
// pseudo-versions above the newest tag, so the tags describe nothing anyone
// consumes. The mechanical reason turned out to be measurable rather than
// cultural, and it is what TestDeclaredVersionIsPinnableByGoModules pins.

// TestDeclaredVersionIsPinnableByGoModules is the whole of #22 in one
// assertion.
//
// The newest tag in this repository is `v2.3.0+rc0`, and Go cannot use it.
// Build metadata - anything after a "+" - is not permitted in a module version
// (only the historical "+incompatible" is), so the module proxy does not see
// that tag as a version at all:
//
//	$ go list -m github.com/footprintai/grandturks-client/v2@v2.3.0+rc0
//	github.com/footprintai/grandturks-client/v2 v2.3.1-0.20241023162104-d0429b3a9b13
//
// It resolves to a PSEUDO-VERSION - Go's way of saying "this commit has no tag
// on the path to one". Which is exactly what grandturks' go.mod contains, and
// exactly what #22 is about: nobody chose pseudo-versions, the newest tag was
// simply unusable and Go fell back.
//
// So the version this module declares must be one a tag can legally carry. A
// release candidate is written `2.4.0-rc.0` - a PRE-RELEASE, which is valid
// and sorts before `2.4.0` - never `2.4.0+rc0`.
func TestDeclaredVersionIsPinnableByGoModules(t *testing.T) {
	declared := GetVersion()

	if strings.Contains(declared, "+") {
		t.Errorf("declared version %q carries build metadata; Go module versions may not, "+
			"so a tag of this shape resolves to a pseudo-version and is invisible to consumers "+
			"(use a pre-release like 2.4.0-rc.0 instead)", declared)
	}
	if _, err := goversion.NewSemver(declared); err != nil {
		t.Errorf("declared version %q is not valid semver: %v", declared, err)
	}
}

// TestTagNameIsTheDeclaredVersion pins the contract the release workflow
// depends on: it compares the pushed tag against this string, and refuses to
// publish a tag the module does not agree it is. Without that check a
// `v2.5.0` tag could ship a binary that reports 2.4.0 - which is the
// "self-reports version:2.3.0+rc0" complaint in #15, one release later.
func TestTagNameIsTheDeclaredVersion(t *testing.T) {
	want := "v" + GetVersion()
	if got := TagName(); got != want {
		t.Errorf("TagName() = %q, want %q", got, want)
	}
	if !strings.HasPrefix(TagName(), "v") {
		t.Errorf("TagName() = %q, want a leading v - git tags in this repo carry one", TagName())
	}
}

// TestGreatThan covers the comparison callers actually make, and one case that
// looks like a comparison and is not: build metadata is ignored by semver
// precedence, so `2.3.0+rc0` and `2.3.0` compare EQUAL. A tag named "+rc0"
// therefore does not sort before its release even where it is understood - a
// second reason not to write one.
func TestGreatThan(t *testing.T) {
	for _, tc := range []struct {
		name string
		v1   string
		v2   string
		want bool
	}{
		{"newer minor", "2.4.0", "2.3.0", true},
		{"older minor", "2.3.0", "2.4.0", false},
		{"equal", "2.4.0", "2.4.0", false},
		{"patch", "2.4.1", "2.4.0", true},
		{"pre-release sorts before its release", "2.4.0-rc.0", "2.4.0", false},
		{"release beats its pre-release", "2.4.0", "2.4.0-rc.0", true},
		{"build metadata is not a pre-release", "2.3.0+rc0", "2.3.0", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := GreatThan(tc.v1, tc.v2); got != tc.want {
				t.Errorf("GreatThan(%q, %q) = %v, want %v", tc.v1, tc.v2, got, tc.want)
			}
		})
	}
}

// TestGetVersionSurvivesAnUnparseableOverride: GetVersion used to dereference
// the result of a parse whose error was discarded, so a version string that
// did not parse was a nil pointer and a panic in `kafeido version`. The
// release procedure now edits this string by hand, which is exactly the
// circumstance in which someone mistypes it.
func TestGetVersionSurvivesAnUnparseableOverride(t *testing.T) {
	original := declaredVersion
	t.Cleanup(func() { declaredVersion = original })

	declaredVersion = "not a version"
	if got := GetVersion(); got != "not a version" {
		t.Errorf("GetVersion() = %q, want the raw string back rather than a panic or an empty value", got)
	}
}

// TestGreatThanWithAnUnparseableVersion: GreatThan discarded both parse
// errors and then dereferenced the results, so any caller comparing a version
// string it did not construct itself - a value from a server, a config file,
// or a release tag someone mistyped - panicked with a nil pointer instead of
// getting an answer.
//
// The answer for input that is not a version is false: an unknown version is
// not greater than anything.
func TestGreatThanWithAnUnparseableVersion(t *testing.T) {
	for _, tc := range []struct {
		name string
		v1   string
		v2   string
	}{
		{"first is garbage", "not a version", "2.0.0"},
		{"second is garbage", "2.0.0", "not a version"},
		{"both are garbage", "not a version", "also not one"},
		{"empty", "", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := GreatThan(tc.v1, tc.v2); got != false {
				t.Errorf("GreatThan(%q, %q) = %v, want false", tc.v1, tc.v2, got)
			}
		})
	}
}
