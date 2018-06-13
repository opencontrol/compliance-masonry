/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package masonry

import (
	"os"
)

// Run the Masonry Command structure
func Run() error {
	cmd := NewMasonryCommand(os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}
