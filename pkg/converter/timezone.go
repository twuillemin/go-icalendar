package converter

import (
	"fmt"
	"sort"

	"github.com/twuillemin/go-icalendar/pkg/entity"
	"github.com/twuillemin/go-icalendar/pkg/parser"
)

// ReadTimeZoneProperty reads a Timezone object
func ReadTimeZone(object parser.Object) (*entity.Timezone, error) {
	if object.Name != "VTIMEZONE" {
		return nil, fmt.Errorf("the object given is of type %s, expected \"VTIMEZONE\"", object.Name)
	}

	id, err := getMandatorySingleString(object, "TZID")
	if err != nil {
		return nil, err
	}

	lastModified, err := getOptionalSingleDateUTC(object, "LAST-MODIFIED")
	if err != nil {
		return nil, err
	}

	url, err := getOptionalSingleString(object, "TZURL")
	if err != nil {
		return nil, err
	}

	properties := make([]entity.TimezoneProperty, 0, len(object.Objects["STANDARD"])+len(object.Objects["DAYLIGHT"]))

	for _, std := range object.Objects["STANDARD"] {
		standard, err := ReadTimeZoneProperty(std)
		if err != nil {
			return nil, err
		}

		properties = append(properties, *standard)
	}

	for _, std := range object.Objects["DAYLIGHT"] {
		standard, err := ReadTimeZoneProperty(std)
		if err != nil {
			return nil, err
		}

		properties = append(properties, *standard)
	}

	sort.Slice(properties, func(i, j int) bool {
		return properties[i].StartLocal.Before(properties[j].StartLocal)
	})

	return &entity.Timezone{
		ID:           id,
		URL:          url,
		LastModified: lastModified,
		Properties:   properties,
	}, nil
}

// ReadTimeZoneProperty reads the property object in a Timezone. The property can be a Standard or a DayLight.
func ReadTimeZoneProperty(object parser.Object) (*entity.TimezoneProperty, error) {
	var tzType entity.TimezonePropertyType

	switch object.Name {
	case "STANDARD":
		tzType = entity.Standard
	case "DAYLIGHT":
		tzType = entity.Daylight
	default:
		return nil, fmt.Errorf("the object given is of type %s, expected \"STANDARD\" or \"DAYLIGHT\"", object.Name)
	}

	startLocal, err := getMandatorySingleDateUTC(object, "DTSTART")
	if err != nil {
		return nil, err
	}

	offsetFrom, err := getMandatorySingleHourMinute(object, "TZOFFSETFROM")
	if err != nil {
		return nil, err
	}

	offsetTo, err := getMandatorySingleHourMinute(object, "TZOFFSETTO")
	if err != nil {
		return nil, err
	}

	rule, err := getOptionalSingleString(object, "RRULE")
	if err != nil {
		return nil, err
	}

	comments := getOptionalMultipleString(object, "COMMENT")
	names := getOptionalMultipleString(object, "TZNAME")

	rDatesLocal, err := getOptionalMultipleDateUTC(object, "RDATE")
	if err != nil {
		return nil, err
	}

	return &entity.TimezoneProperty{
		Type:        tzType,
		StartLocal:  startLocal,
		OffsetFrom:  offsetFrom,
		OffsetTo:    offsetTo,
		Rule:        rule,
		Comments:    comments,
		RDatesLocal: rDatesLocal,
		Names:       names,
	}, nil
}
