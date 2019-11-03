package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twuillemin/go-icalendar/pkg/parser"
)

func TestReadCalendarEventEmpty(t *testing.T) {
	_, err := ReadEvent(getEventEmptyObject()[0])
	assert.NoError(t, err)
}

func getEventEmptyObject() []parser.Object {
	return []parser.Object{
		{
			Name: "VEVENT",
		},
	}
}
