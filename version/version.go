/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package version

import (
	"fmt"
	"os"
)

// Version is the version of the build.
// Build details
var (
	Version string
	Commit  string
	Date    string
)

// PrintVersion returns the version for the command version and --version flag
func PrintVersion() {
	fmt.Printf("masonry version: %s, build: %s, date: %s\n", Version, Commit, Date)
	os.Exit(0)
}
