package common

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	filename := "testdata/testfile.txt"
	expected := "This is my voice in a test!"
	result := string(ParseFile(filename))

	if expected != result {
		t.Errorf("Expected %s to equal %s", result, expected)
	}
}
