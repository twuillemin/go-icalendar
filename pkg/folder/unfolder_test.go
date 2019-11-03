package folder

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnfoldLines(t *testing.T) {
	raw := "This is a lo\n" +
		" ng description \n" +
		"\tthat exists on a long line.\n" +
		"This is a short line\n" +
		"This is another short line"

	expected := []string{
		"This is a long description that exists on a long line.",
		"This is a short line",
		"This is another short line",
	}

	testUnfold(t, raw, expected)
}

func TestCleanWhiteLines(t *testing.T) {
	raw := "line1\n\nline2"

	expected := []string{
		"line1",
		"line2",
	}

	testUnfold(t, raw, expected)
}

func testUnfold(t *testing.T, raw string, expectedLines []string) {
	unfoldedLines, err := UnfoldLines(strings.NewReader(raw))
	assert.NoError(t, err)

	expectedLength := len(expectedLines)
	readLength := len(unfoldedLines)
	assert.Equal(t, expectedLength, readLength, "expected %d lines, but %d read", expectedLength, readLength)

	for i, expectedLine := range expectedLines {
		readLine := unfoldedLines[i]
		assert.Equal(t, expectedLine, readLine, "difference detected at line %d", i)
	}
}
