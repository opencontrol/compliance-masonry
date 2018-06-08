package main

import (
	"os"

	"github.com/opencontrol/compliance-masonry/pkg/cmd/masonry"
)

func main() {
	if err := masonry.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
