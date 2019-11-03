package converter

import (
	"fmt"

	"github.com/twuillemin/go-icalendar/pkg/entity"
	"github.com/twuillemin/go-icalendar/pkg/parser"
)

// ReadEvent reads an Event object
func ReadEvent(object parser.Object) (*entity.Event, error) {
	if object.Name != "VEVENT" {
		return nil, fmt.Errorf("the object given is of type %s, expected \"VEVENT\"", object.Name)
	}

	return &entity.Event{}, nil
}
