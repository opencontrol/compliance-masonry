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
	"strings"
)

func main() {
	var binary string
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		binary = "masonry.exe"
	} else {
		binary = "masonry"
	}

	envPaths := os.Getenv("PATH")
	basepath := filepath.Dir(os.Args[0])
	prog, err := filepath.Abs(filepath.Join(basepath, binary))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(prog); os.IsNotExist(err) {
		for _, envp := range strings.Split(envPaths, ":") {
			cprog, err := filepath.Abs(filepath.Join(envp, binary))
			if _, err := os.Stat(cprog); err == nil {
				prog = filepath.Join(envp, binary)
			}
			if err != nil {
				log.Fatal(err)
			}
		}
	}
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
		if len(output) == 0 {
			log.Fatal("Error: cannot find the 'masonry' executable")
		}
	}
	fmt.Printf("%s\n", output)
}
