package parser

import (
	"fmt"
	"strings"
)

type Attribute struct {
	Name  string
	Value string
	Info  map[string]string
}

// readAttributeLine reads an attribute line. The function returns the name of the attribute, its Value and a map of
// meta-information.
func readAttributeLine(attributeLine string) (*Attribute, error) {
	items, err := readLineParts(attributeLine)
	if err != nil {
		return nil, err
	}

	if len(items) < 2 {
		return nil, fmt.Errorf("an attribute line is expected to have at least 2 items")
	}

	name := items[0]
	value := items[len(items)-1]
	info := make(map[string]string)

	if len(name) == 0 {
		return nil, fmt.Errorf("an attribute line is expected to begin with a name")
	}

	for i := 1; i < len(items)-1; i++ {
		firstEqualIndex := strings.Index(items[i], "=")
		if firstEqualIndex == -1 {
			info[items[i]] = ""
		} else {
			parameterName := items[i][:firstEqualIndex]
			parameterValue := items[i][firstEqualIndex+1:]
			info[parameterName] = parameterValue
		}
	}

	attribute := &Attribute{
		Name:  name,
		Value: value,
		Info:  info,
	}

	return attribute, nil
}

// readLineParts reads all the parts from an attribute line. The function return an array of strings, each string being
// a part.
func readLineParts(attribute string) ([]string, error) {
	items := make([]string, 0, 2)
	toRead := attribute

	for len(toRead) > 0 {
		item, lastIndexRead, err := readNextLinePart(toRead)
		if err != nil {
			return nil, err
		}

		items = append(items, item)

		if toRead[lastIndexRead] == ':' {
			items = append(items, toRead[lastIndexRead+1:])
			toRead = ""
		} else {
			toRead = toRead[lastIndexRead+1:]
		}
	}

	return items, nil
}

// readNextLinePart reads the next part in an attribute string. Item separators are ; (for multiple item), : (separating
// parameters from Value) or simply end of line. The readNextLinePart function takes care of Value between double
// quotes. The function returns the part read (without double quotes) and the index of the last character read.
func readNextLinePart(str string) (string, int, error) {
	inQuote := false
	accumulator := ""

	for i, c := range str {
		switch c {
		case '"':
			if inQuote {
				if str[i+1] != ';' && str[i+1] != ':' {
					return "", -1, fmt.Errorf("after a closing double quote a colon or a semi-colon is expected")
				}
			}

			inQuote = !inQuote

		case ';', ':':
			if inQuote {
				accumulator += string(c)
			} else {
				return accumulator, i, nil
			}

		default:
			accumulator += string(c)
		}
	}

	return accumulator, len(str) - 1, nil
}
