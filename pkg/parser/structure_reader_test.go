package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEmptyLines(t *testing.T) {
	var raw []string

	elements, err := readStructure(raw)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(elements))
}

func TestReadSimpleEmptyObject(t *testing.T) {
	raw := []string{
		"BEGIN:VCALENDAR",
		"END:VCALENDAR",
	}

	elements, err := readStructure(raw)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(elements))

	calendar := elements[0]
	assert.Equal(t, "VCALENDAR", calendar.name)
	assert.Equal(t, 0, len(calendar.attributes))
	assert.Equal(t, 0, len(calendar.structureElements))
}

func TestReadObjectWithNoName(t *testing.T) {
	raw := []string{
		"BEGIN:",
		"VERSION:2.0",
		"PRODID:-//hacksw/handcal//NONSGML v1.0//EN",
		"END:",
	}

	elements, err := readStructure(raw)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(elements))

	element := elements[0]
	assert.Equal(t, "", element.name)
	assert.Equal(t, 2, len(element.attributes))
	assert.Equal(t, 0, len(element.structureElements))
}

func TestReadElementInElement(t *testing.T) {
	raw := []string{
		"BEGIN:VCALENDAR",
		"VERSION:2.0",
		"PRODID:-//hacksw/handcal//NONSGML v1.0//EN",
		"BEGIN:VEVENT",
		"SUMMARY:Bastille Day Party",
		"END:VEVENT",
		"END:VCALENDAR",
	}

	elements, err := readStructure(raw)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(elements))

	calendar := elements[0]
	assert.Equal(t, "VCALENDAR", calendar.name)
	assert.Equal(t, 2, len(calendar.attributes))
	assert.Equal(t, 1, len(calendar.structureElements))

	event := calendar.structureElements[0]
	assert.Equal(t, "VEVENT", event.name)
	assert.Equal(t, 1, len(event.attributes))
	assert.Equal(t, 0, len(event.structureElements))
}

func TestRaiseErrorOnUnmatchedElement(t *testing.T) {
	raw := []string{
		"BEGIN:VCALENDAR",
		"VERSION:2.0",
		"PRODID:-//hacksw/handcal//NONSGML v1.0//EN",
	}

	_, err := readStructure(raw)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "end"))
}

func TestRaiseErrorOnMismatchedElement(t *testing.T) {
	raw := []string{
		"BEGIN:VCALENDAR",
		"BEGIN:VEVENT",
		"END:VCALENDAR",
		"END:VEVENT",
	}

	_, err := readStructure(raw)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "mismatch"))
}

func TestRaiseErrorOnRootAttributes(t *testing.T) {
	raw := []string{
		"VERSION:2.0",
		"BEGIN:VCALENDAR",
		"PRODID:-//hacksw/handcal//NONSGML v1.0//EN",
		"END:VCALENDAR",
	}

	_, err := readStructure(raw)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "level"))
}
