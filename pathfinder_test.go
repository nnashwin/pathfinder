package pathfinder

import (
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
