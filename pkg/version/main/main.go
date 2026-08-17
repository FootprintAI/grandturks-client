package main

import (
	"fmt"

	"github.com/footprintai/grandturks-client/v2/pkg/version"
)

// Prints the tag this tree expects to be released under. The release
// workflow runs this and refuses to publish a tag that disagrees, so the
// output format is a contract - see pkg/version.TagName.
func main() {
	fmt.Printf("%s", version.TagName())
}
