package converter

import (
	"fmt"

	"github.com/twuillemin/go-icalendar/pkg/entity"
	"github.com/twuillemin/go-icalendar/pkg/parser"
)

// ReadTimeZoneProperty reads a Calendar object
func ReadCalendar(object parser.Object) (*entity.Calendar, error) {
	if object.Name != "VCALENDAR" {
		return nil, fmt.Errorf("the object given is of type %s", object.Name)
	}

	if len(object.Attributes["VERSION"]) != 1 {
		return nil, fmt.Errorf("the attribute \"VERSION\" is mandatory")
	}

	version := object.Attributes["VERSION"][0].Value
	if version != "2.0" {
		return nil, fmt.Errorf("the version of the calendar was %s, \"2.0\" is mandatory", version)
	}

	if len(object.Attributes["PRODID"]) != 1 {
		return nil, fmt.Errorf("the attribute \"PRODID\" is mandatory")
	}

	productID := object.Attributes["PRODID"][0].Value

	events := make([]entity.Event, 0)

	for _, subObject := range object.Objects["VEVENT"] {
		event, err := ReadEvent(subObject)
		if err != nil {
			return nil, err
		}

		events = append(events, *event)
	}

	return &entity.Calendar{
		ProductID: productID,
		Version:   version,
		Events:    events,
	}, nil
}
