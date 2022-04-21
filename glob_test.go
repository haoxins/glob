package glob

import (
	"fmt"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGlob(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Glob Suite")
}

var _ = Describe("Test Glob", func() {
	run("mkdir", "-p", "temp/a/b/c", "temp/b/c/d")
	run("touch", "temp/a/a.go", "temp/a/b/b.go", "temp/a/b/c/c.go")
	run("touch", "temp/b/b.go", "temp/b/c/c.go", "temp/b/c/d/d.go")

	It("Glob should work", func() {
		matches, err := Glob(".", "**/*.go")
		Expect(err).To(BeNil())
		Expect(len(matches)).To(Equal(8))

		fmt.Println(matches)

		run("rm", "-r", "temp")

		matches, err = Glob(".", "**/*.md")
		Expect(err).To(BeNil())
		Expect(len(matches)).To(Equal(1))
		Expect(matches[0]).To(Equal("Readme.md"))

		matches, err = Glob(".", "*.md")
		Expect(err).To(BeNil())
		Expect(len(matches)).To(Equal(1))
		Expect(matches[0]).To(Equal("Readme.md"))

		matches, err = Glob(".", "**/*.yaml")
		Expect(err).To(BeNil())
		Expect(len(matches)).To(Equal(1))
		Expect(matches[0]).To(Equal(".github/workflows/test.yaml"))
	})
})

func run(name string, arg ...string) (out string, err error) {
	data, err := exec.Command(name, arg...).Output()
	if err != nil {
		return err.Error(), err
	}

	return string(data[:]), nil
}
