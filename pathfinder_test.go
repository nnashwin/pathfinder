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

func TestCreateFile(t *testing.T) {
	cases := []struct {
		path    string
		failMsg string
	}{
		{
			"filedir/nestedDir2/testFile",
			"The path created was incorrect.",
		},
		{
			"filedir.file",
			"The create file without a nested path failed",
		},
	}

	// change the dir to the fixture dir before running the tests
	err := os.Chdir("./fixtures")
	if err != nil {
		t.Error(err)
	}

	for _, c := range cases {
		if DoesExist(c.path) != false {
			t.Error("Cleanup on the last test was not done correct.  The path should be false before we create the file.")
		}

		err = CreateFile(c.path)
		if err != nil {
			t.Errorf("The following problem occurred when running CreateFile: %s", err)
		}

		if DoesExist(c.path) != true {
			t.Error("The Create File function didn't create a path.")
		}

		err = os.RemoveAll(c.path)
		if err != nil {
			t.Errorf("Cleanup in the CreatePath test failed with the following error: %s", err)
		}
	}

	err = os.Chdir("../")
	if err != nil {
		t.Error(err)
	}
}

func TestCreateDir(t *testing.T) {
	cases := []struct {
		path    string
		failMsg string
	}{
		{
			"./fixtures/filedir/seconddir/",
			"The path created was incorrect.",
		},
	}

	for _, c := range cases {
		if DoesExist(c.path) != false {
			t.Error("Cleanup on the last test was not done correct.  The path should be false before we create the file.")
		}

		err := CreateDir(c.path)
		if err != nil {
			t.Errorf("The following problem occurred when running CreateDir: %s", err)
		}

		if DoesExist(c.path) != true {
			t.Error("The Create Dir function didn't create a path.")
		}
	}

	err := os.RemoveAll("./fixtures/filedir")
	if err != nil {
		t.Errorf("Cleanup in the CreateDir test failed with the following error: %s", err)
	}
}

func TestGetDir(t *testing.T) {
	cases := []struct {
		path     string
		expected string
	}{
		{
			"/.dir/dir/dir.file",
			"/.dir/dir/",
		},
		{
			"dir.file",
			"",
		},
		{
			"/../dir.file",
			"/",
		},
	}

	for _, c := range cases {
		if GetDir(c.path) != c.expected {
			t.Errorf("The Get Dir function didn't return the expected dir from the path: %s", c.path)
		}
	}
}
