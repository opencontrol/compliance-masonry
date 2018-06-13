/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	var binary string
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		binary = "masonry.exe"
	} else {
		binary = "masonry"
	}

	prog, err := filepath.Abs(filepath.Join(binary))
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args
	if len(args) > 1 {
		cmd = exec.Command(prog, args[1:]...)
	} else {
		cmd = exec.Command(prog)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		//do nothing
	}
	fmt.Printf("%s\n", output)
}
