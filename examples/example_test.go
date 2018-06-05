package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common/mocks"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var _ = Describe("Exampleplugin", func() {
	// example of unit test with mocks.
	Describe("simpleDataExtract", func() {
		Context("When there is no data for the standard-control combo", func() {
			It("should print 'no data'", func() {
				// create mock workspace
				ws := new(mocks.Workspace)
				ws.On("GetAllVerificationsWith", "NIST-800-53", "CM-2").Return(common.Verifications{})
				// test function expecting "no data"
				p := plugin{ws}
				data := simpleDataExtract(p)
				Expect(data).To(Equal("no data"))
			})
		})
		Context("When there is data for the standard-control combo", func() {
			It("should print the 'IMPLEMENTED', given that's expected", func() {
				// create mock workspace
				ws := new(mocks.Workspace)
				satisfies := new(mocks.Satisfies)
				satisfies.On("GetImplementationStatus").Return("IMPLEMENTED")
				ws.On("GetAllVerificationsWith", "NIST-800-53", "CM-2").Return(
					common.Verifications{common.Verification{SatisfiesData: satisfies}})
				// test function expecting "IMPLEMENTED"
				p := plugin{ws}
				data := simpleDataExtract(p)
				Expect(data).To(Equal("IMPLEMENTED"))
				assert.Equal(GinkgoT(), data, "IMPLEMENTED")
			})
		})
	})
	// Example of reading the standard output.
	Describe("run", func() {
		Context("When running it on data in a workspace", func() {
			It("should find the data and print it out to standard out", func() {
				wsPath := filepath.Join("..", "fixtures", "opencontrol_fixtures")
				certPath := filepath.Join("..", "..", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml")
				buffer := NewBuffer()
				run(wsPath, certPath, buffer)
				Expect(buffer).To(Say("partial"))
			})
		})
	})
	// Example of Running Masonry (with "get") and then the Plugin executables.
	Describe("running the executable", func() {
		BeforeEach(func() {
			cleanupOpencontrolWorkspace()
		})
		Context("when running the executable", func() {
			It("should build and run with two arguments", func() {
				masonry := Masonry("--verbose", "get")
				Eventually(masonry).Should(Exit(0))
				p := Plugin(filepath.Join("opencontrols"), filepath.Join("opencontrols", "certifications"))
				Eventually(p).Should(Exit(0))
				// Should match the implementation status of CloudFormation in aws-compliance, which is
				// "none"
				Eventually(p.Out.Contents()).Should(ContainSubstring("none"))
			})
		})
		AfterEach(func() {
			cleanupOpencontrolWorkspace()
			CleanupBuildArtifacts()
		})
	})
})

func cleanupOpencontrolWorkspace() {
	os.RemoveAll("opencontrols")
}

func Masonry(args ...string) *Session {
	path, err := Build("github.com/opencontrol/compliance-masonry")
	Expect(err).NotTo(HaveOccurred())
	return createCommand(path, args...)
}
func Plugin(args ...string) *Session {
	path, err := Build("github.com/opencontrol/compliance-masonry/exampleplugin")
	Expect(err).NotTo(HaveOccurred())
	return createCommand(path, args...)
}

func createCommand(cmdPath string, args ...string) *Session {
	cmd := exec.Command(cmdPath, args...)
	session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
	session.Wait(20 * time.Second)
	Expect(err).NotTo(HaveOccurred())
	return session
}
