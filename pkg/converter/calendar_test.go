package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twuillemin/go-icalendar/pkg/parser"
)

func TestReadCalendarEmpty(t *testing.T) {
	calendar, err := ReadCalendar(getCalendarEmptyObject()[0])
	assert.NoError(t, err)
	assert.Equal(t, "2.0", calendar.Version)
	assert.Equal(t, "Test Prod", calendar.ProductID)
}

func getCalendarEmptyObject() []parser.Object {
	return []parser.Object{
		{
			Name: "VCALENDAR",
			Attributes: map[string][]parser.Attribute{
				"VERSION": {{Name: "VERSION", Value: "2.0"}},
				"PRODID":  {{Name: "PRODID", Value: "Test Prod"}},
			},
		},
	}
}
