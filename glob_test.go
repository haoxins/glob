package glob

import "os/exec"
import "testing"

func TestGlob(t *testing.T) {
	Run("mkdir", "-p", "temp/a/b/c", "temp/b/c/d")
	Run("touch", "temp/a/a.go", "temp/a/b/b.go", "temp/a/b/c/c.go")
	Run("touch", "temp/b/b.go", "temp/b/c/c.go", "temp/b/c/d/d.go")

	matches, err := Glob(".", "**/*.go")
	ifError(t, err)
	shouldEqual(t, len(matches) == 8, "matches length is incorrect")

	t.Log(matches)

	Run("rm", "-r", "temp")

	matches, err = Glob(".", "**/*.md")
	ifError(t, err)
	shouldEqual(t, len(matches) == 1, "matches length is incorrect")
	shouldEqual(t, matches[0] == "Readme.md", "matches is incorrect")

	matches, err = Glob(".", "*.md")
	ifError(t, err)
	shouldEqual(t, len(matches) == 1, "matches length is incorrect")
	shouldEqual(t, matches[0] == "Readme.md", "matches is incorrect")

	matches, err = Glob(".", ".*.*")
	ifError(t, err)
	shouldEqual(t, len(matches) == 1, "matches length is incorrect")
	shouldEqual(t, matches[0] == ".travis.yml", "matches is incorrect")
}

func Run(name string, arg ...string) (out string, err error) {
	data, err := exec.Command(name, arg...).Output()
	if err != nil {
		return err.Error(), err
	}

	return string(data[:]), nil
}

func ifError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func shouldEqual(t *testing.T, equal bool, msg string) {
	if !equal {
		t.Fatal(msg)
	}
}
