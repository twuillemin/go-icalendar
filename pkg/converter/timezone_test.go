package converter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/twuillemin/go-icalendar/pkg/entity"
	"github.com/twuillemin/go-icalendar/pkg/parser"
)

func TestReadTimeZone(t *testing.T) {
	obj := parser.Object{
		Name: "VTIMEZONE",
		Objects: map[string][]parser.Object{
			"DAYLIGHT": getNewYorkDayLightObject(),
			"STANDARD": getNewYorkStandardObject(),
		},
		Attributes: map[string][]parser.Attribute{
			"TZID":          {{Name: "TZID", Value: "America/New_York"}},
			"LAST-MODIFIED": {{Name: "LAST-MODIFIED", Value: "20050809T050000Z"}},
		},
	}

	timezone, err := ReadTimeZone(obj)
	assert.NoError(t, err)
	assert.Equal(t, 7, len(timezone.Properties))
}

func TestReadTimeZonePropertyDayLight(t *testing.T) {
	obj := getNewYorkDayLightObject()[0]

	dtStart, _ := time.Parse(time.RFC3339, "1967-04-30T02:00:00Z")

	prop, err := ReadTimeZoneProperty(obj)
	assert.NoError(t, err)
	assert.Equal(t, entity.Daylight, prop.Type)
	assert.Equal(t, dtStart, prop.StartLocal)
	assert.Equal(t, "FREQ=YEARLY;BYMONTH=4;BYDAY=-1SU;UNTIL=19730429T070000Z", prop.Rule)
	assert.Equal(t, entity.HourMinute{Hour: -5, Minute: 0}, prop.OffsetFrom)
	assert.Equal(t, entity.HourMinute{Hour: -4, Minute: 0}, prop.OffsetTo)
	assert.Equal(t, "EDT", prop.Names[0])
}

func TestReadTimeZonePropertyStandard(t *testing.T) {
	obj := getNewYorkStandardObject()[0]

	dtStart, _ := time.Parse(time.RFC3339, "1967-10-29T02:00:00Z")

	prop, err := ReadTimeZoneProperty(obj)
	assert.NoError(t, err)
	assert.Equal(t, entity.Standard, prop.Type)
	assert.Equal(t, dtStart, prop.StartLocal)
	assert.Equal(t, "FREQ=YEARLY;BYMONTH=10;BYDAY=-1SU;UNTIL=20061029T060000Z", prop.Rule)
	assert.Equal(t, entity.HourMinute{Hour: -4, Minute: 0}, prop.OffsetFrom)
	assert.Equal(t, entity.HourMinute{Hour: -5, Minute: 0}, prop.OffsetTo)
	assert.Equal(t, "EST", prop.Names[0])
}

func getNewYorkDayLightObject() []parser.Object {
	return []parser.Object{
		{
			Name: "DAYLIGHT",
			Attributes: map[string][]parser.Attribute{
				"DTSTART":      {{Name: "DTSTART", Value: "19670430T020000"}},
				"RRULE":        {{Name: "RRULE", Value: "FREQ=YEARLY;BYMONTH=4;BYDAY=-1SU;UNTIL=19730429T070000Z"}},
				"TZOFFSETFROM": {{Name: "TZOFFSETFROM", Value: "-0500"}},
				"TZOFFSETTO":   {{Name: "TZOFFSETTO", Value: "-0400"}},
				"TZNAME":       {{Name: "TZNAME", Value: "EDT"}},
			},
		},
		{
			Name: "DAYLIGHT",
			Attributes: map[string][]parser.Attribute{
				"DTSTART":      {{Name: "DTSTART", Value: "19740106T020000"}},
				"RDATE":        {{Name: "RDATE", Value: "19750223T020000"}},
				"TZOFFSETFROM": {{Name: "TZOFFSETFROM", Value: "-0500"}},
				"TZOFFSETTO":   {{Name: "TZOFFSETTO", Value: "-0400"}},
				"TZNAME":       {{Name: "TZNAME", Value: "EDT"}},
			},
		},
		{
			Name: "DAYLIGHT",
			Attributes: map[string][]parser.Attribute{
				"DTSTART":      {{Name: "DTSTART", Value: "19760425T020000"}},
				"RRULE":        {{Name: "RRULE", Value: "FREQ=YEARLY;BYMONTH=4;BYDAY=-1SU;UNTIL=19860427T070000Z"}},
				"TZOFFSETFROM": {{Name: "TZOFFSETFROM", Value: "-0500"}},
				"TZOFFSETTO":   {{Name: "TZOFFSETTO", Value: "-0400"}},
				"TZNAME":       {{Name: "TZNAME", Value: "EDT"}},
			},
		},
		{
			Name: "DAYLIGHT",
			Attributes: map[string][]parser.Attribute{
				"DTSTART":      {{Name: "DTSTART", Value: "19870405T020000"}},
				"RRULE":        {{Name: "RRULE", Value: "FREQ=YEARLY;BYMONTH=4;BYDAY=1SU;UNTIL=20060402T070000Z"}},
				"TZOFFSETFROM": {{Name: "TZOFFSETFROM", Value: "-0500"}},
				"TZOFFSETTO":   {{Name: "TZOFFSETTO", Value: "-0400"}},
				"TZNAME":       {{Name: "TZNAME", Value: "EDT"}},
			},
		},
		{
			Name: "DAYLIGHT",
			Attributes: map[string][]parser.Attribute{
				"DTSTART":      {{Name: "DTSTART", Value: "20070311T020000"}},
				"RRULE":        {{Name: "RRULE", Value: "FREQ=YEARLY;BYMONTH=3;BYDAY=2SU"}},
				"TZOFFSETFROM": {{Name: "TZOFFSETFROM", Value: "-0500"}},
				"TZOFFSETTO":   {{Name: "TZOFFSETTO", Value: "-0400"}},
				"TZNAME":       {{Name: "TZNAME", Value: "EDT"}},
			},
		},
	}
}

func getNewYorkStandardObject() []parser.Object {
	return []parser.Object{
		{
			Name: "STANDARD",
			Attributes: map[string][]parser.Attribute{
				"DTSTART":      {{Name: "DTSTART", Value: "19671029T020000"}},
				"RRULE":        {{Name: "RRULE", Value: "FREQ=YEARLY;BYMONTH=10;BYDAY=-1SU;UNTIL=20061029T060000Z"}},
				"TZOFFSETFROM": {{Name: "TZOFFSETFROM", Value: "-0400"}},
				"TZOFFSETTO":   {{Name: "TZOFFSETTO", Value: "-0500"}},
				"TZNAME":       {{Name: "TZNAME", Value: "EST"}},
			},
		},
		{
			Name: "STANDARD",
			Attributes: map[string][]parser.Attribute{
				"DTSTART":      {{Name: "DTSTART", Value: "20071104T020000"}},
				"RRULE":        {{Name: "RRULE", Value: "FREQ=YEARLY;BYMONTH=11;BYDAY=1SU"}},
				"TZOFFSETFROM": {{Name: "TZOFFSETFROM", Value: "-0500"}},
				"TZOFFSETTO":   {{Name: "TZOFFSETTO", Value: "-0400"}},
				"TZNAME":       {{Name: "TZNAME", Value: "EST"}},
			},
		},
	}
}
