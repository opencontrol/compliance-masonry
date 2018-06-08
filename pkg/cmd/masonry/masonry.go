package masonry

import (
	"os"
)

// Run the Masonry Command structure
func Run() error {
	cmd := NewMasonryCommand(os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}
