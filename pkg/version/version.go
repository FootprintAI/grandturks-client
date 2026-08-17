package version

import (
	"fmt"
	"runtime/debug"

	goversion "github.com/hashicorp/go-version"
)

var (
	BuildTime = ""

	// declaredVersion is the version the NEXT tag cut from this tree will
	// carry, and the single source of truth for what this module calls
	// itself. The release workflow refuses to publish a tag that disagrees
	// with it (see RELEASING.md), which is what stops a v2.5.0 release
	// shipping a binary that reports something else.
	//
	// It must stay a legal Go module version. It read "2.3.0+rc0" until
	// grandturks-client#22: build metadata is not permitted in a module
	// version, so the matching tag was invisible to the proxy and consumers
	// fell back to pseudo-versions. A release candidate is "2.4.0-rc.0".
	declaredVersion = "2.4.0"
)

func GreatThan(v1, v2 string) bool {
	v1v, _ := goversion.NewVersion(v1)
	v2v, _ := goversion.NewVersion(v2)

	return v1v.GreaterThan(v2v)
}

// GetVersion returns the declared version, normalised if it parses.
//
// A string that does not parse is returned as-is rather than dereferenced:
// this value is edited by hand at release time, and a mistyped one used to be
// a nil pointer and a panic in `kafeido version`.
func GetVersion() string {
	v, err := goversion.NewVersion(declaredVersion)
	if err != nil {
		return declaredVersion
	}
	return v.String()
}

// TagName is the git tag that corresponds to this version.
//
// The release workflow compares the pushed tag against this, so tag and module
// cannot drift apart the way they have (grandturks-client#22).
func TagName() string {
	return "v" + GetVersion()
}

func GetBuildTime() string {
	return BuildTime
}

func GetCommitHash() string {
	info, _ := debug.ReadBuildInfo()
	var rev string = "<none>"
	var dirty string = ""
	for _, v := range info.Settings {
		if v.Key == "vcs.revision" {
			rev = v.Value
		}
		if v.Key == "vcs.modified" {
			if v.Value == "true" {
				dirty = "-dirty"
			} else {
				dirty = ""
			}
		}
	}
	return rev + dirty
}

func Print() {
	fmt.Printf("version:%s, build time:%s\nhashid:%s\n", GetVersion(), GetBuildTime(), GetCommitHash())
}
