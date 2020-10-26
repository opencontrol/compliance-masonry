/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package masonry_test

import (
	"bufio"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

// Masonry is used to launch tests on the CLI
func Masonry(args ...string) *gexec.Session {
	RegisterFailHandler(Fail)
	path, err := gexec.Build("github.com/opencontrol/compliance-masonry/cmd/masonry")
	Expect(err).NotTo(HaveOccurred())
	cmd := exec.Command(path, args...)
	stdin, err := cmd.StdinPipe()
	Expect(err).ToNot(HaveOccurred())
	buffer := bufio.NewWriter(stdin)
	_, _ = buffer.WriteString(strings.Join(args, " "))
	_ = buffer.Flush()
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	session.Wait()
	gexec.CleanupBuildArtifacts()
	return session
}
