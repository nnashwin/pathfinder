package pathfinder

import (
	"os"
	"testing"
)

func TestDoesExist(t *testing.T) {
	cases := []struct {
		path     string
		expected bool
		failMsg  string
	}{
		{
			"./fixtures/fail.txt",
			false,
			"fail.txt should not be found.",
		},
		{
			"./fixtures/pass.txt",
			true,
			"the pass.txt fixture should have been found.",
		},
	}

	for _, c := range cases {
		actual := DoesExist(c.path)
		if actual != c.expected {
			t.Error(c.failMsg)
		}
	}
}

func TestCreatePath(t *testing.T) {
	cases := []struct {
		path    string
		failMsg string
	}{
		{
			"./fixtures/filedir/nestedDir2/testFile",
			"The path created was incorrect.",
		},

		{
			"./fixtures/filedir/seconddir/",
			"The path created was incorrect.",
		},
	}

	for _, c := range cases {
		if DoesExist(c.path) != false {
			t.Error("Cleanup on the last test was not done correct.  The path should be false before we create the file.")
		}

		err := CreatePath(c.path)
		if err != nil {
			t.Errorf("The following problem occurred when running CreatePath: %s", err)
		}

		if DoesExist(c.path) != true {
			t.Error("The Create Path function didn't create a path.")
		}
	}

	err := os.RemoveAll("./fixtures/filedir")
	if err != nil {
		t.Errorf("Cleanup in the CreatePath test failed with the following error: %s", err)
	}
}
