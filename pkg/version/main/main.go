package main

import (
	"fmt"

	"github.com/footprintai/grandturks-client/v2/pkg/version"
)

func main() {
	fmt.Printf("v%s", version.GetVersion())
}
