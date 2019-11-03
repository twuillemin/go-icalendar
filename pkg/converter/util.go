package converter

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/twuillemin/go-icalendar/pkg/entity"
	"github.com/twuillemin/go-icalendar/pkg/parser"
)

func getOptionalMultipleString(object parser.Object, attributeName string) []string {
	attrs := object.Attributes[attributeName]

	result := make([]string, 0, len(attrs))
	for _, attr := range attrs {
		result = append(result, attr.Value)
	}

	return result
}

func getOptionalMultipleDateUTC(object parser.Object, attributeName string) ([]time.Time, error) {
	attrs := object.Attributes[attributeName]
	result := make([]time.Time, 0, len(attrs))

	for _, attr := range attrs {
		date, err := readDateUTC(attr.Value)
		if err != nil {
			return nil, err
		}

		result = append(result, date)
	}

	return result, nil
}

func getOptionalSingleDateUTC(object parser.Object, attributeName string) (time.Time, error) {
	attr, err := getOptionalSingleAttribute(object, attributeName)
	if err != nil || attr == nil {
		return time.Time{}, err
	}

	return readDateUTC(attr.Value)
}

func getOptionalSingleString(object parser.Object, attributeName string) (string, error) {
	attr, err := getOptionalSingleAttribute(object, attributeName)
	if err != nil || attr == nil {
		return "", err
	}

	return attr.Value, nil
}

func getMandatorySingleDateUTC(object parser.Object, attributeName string) (time.Time, error) {
	attr, err := getMandatorySingleAttribute(object, attributeName)
	if err != nil {
		return time.Time{}, err
	}

	if len(attr.Value) == 0 {
		return time.Time{}, fmt.Errorf("the attribute \"%s\" is mandatory", attributeName)
	}

	return readDateUTC(attr.Value)
}

func getMandatorySingleString(object parser.Object, attributeName string) (string, error) {
	attr, err := getMandatorySingleAttribute(object, attributeName)
	if err != nil {
		return "", fmt.Errorf("the attribute \"%s\" is mandatory", attributeName)
	}

	return attr.Value, nil
}

func getMandatorySingleHourMinute(object parser.Object, attributeName string) (entity.HourMinute, error) {
	attr, err := getMandatorySingleAttribute(object, attributeName)
	if err != nil {
		return entity.HourMinute{}, fmt.Errorf("the attribute \"%s\" is mandatory", attributeName)
	}

	return readHourMinute(attr.Value)
}

func getMandatorySingleAttribute(object parser.Object, attributeName string) (*parser.Attribute, error) {
	attributes := object.Attributes[attributeName]
	switch len(attributes) {
	case 0:
		return nil, fmt.Errorf("the attribute \"%s\" is mandatory", attributeName)
	case 1:
		return &attributes[0], nil
	default:
		return nil, fmt.Errorf("the attribute \"%s\" must only be present one time", attributeName)
	}
}

func getOptionalSingleAttribute(object parser.Object, attributeName string) (*parser.Attribute, error) {
	attributes := object.Attributes[attributeName]
	switch len(attributes) {
	case 0:
		return nil, nil
	case 1:
		return &attributes[0], nil
	default:
		return nil, fmt.Errorf("the attribute \"%s\" must only be present one time", attributeName)
	}
}

func readDateUTC(str string) (time.Time, error) {
	switch len(str) {
	case 0:
		return time.Time{}, nil
	case 15:
		return time.Parse("20060102T150405", str)
	case 16:
		return time.Parse("20060102T150405Z", str)
	default:
		return time.Time{}, fmt.Errorf("expected format date format 20060102T150405(Z)")
	}
}

func readHourMinute(str string) (entity.HourMinute, error) {
	re := regexp.MustCompile(`^\s*([-+]?)(\d{2})(\d{2})\s*$`)
	matched := re.FindStringSubmatch(str)

	hour, err := strconv.Atoi(matched[2])
	if err != nil {
		return entity.HourMinute{}, err
	}

	if matched[1] == "-" {
		hour = -hour
	}

	minute, err := strconv.Atoi(matched[3])
	if err != nil {
		return entity.HourMinute{}, err
	}

	return entity.HourMinute{
		Hour:   hour,
		Minute: minute,
	}, nil
}
