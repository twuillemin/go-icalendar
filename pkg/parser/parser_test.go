package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	raw := []string{
		"BEGIN:VCALENDAR",
		"VERSION:2.0",
		"PRODID:-//hacksw/handcal//NONSGML v1.0//EN",
		"BEGIN:VEVENT",
		"SUMMARY:Bastille Day Party",
		"END:VEVENT",
		"END:VCALENDAR",
		"BEGIN:VCALENDAR",
		"VERSION:3.0",
		"PRODID:-//hacksw/handcal//NONSGML v1.0//EN",
		"BEGIN:VEVENT",
		"SUMMARY:Bastille Day Party Hang over",
		"END:VEVENT",
		"END:VCALENDAR",
	}

	objects, err := Parse(raw)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(objects))
	assert.Equal(t, "Bastille Day Party", objects[0].Objects["VEVENT"][0].Attributes["SUMMARY"][0].Value)
	assert.Equal(t, "Bastille Day Party Hang over", objects[1].Objects["VEVENT"][0].Attributes["SUMMARY"][0].Value)
}
