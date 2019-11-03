package parser

import (
	"fmt"
	"strings"
)

type structureElement struct {
	name              string
	structureElements []*structureElement
	attributes        []string
}

type readingAccumulator struct {
	lastLineIndex     int
	structureElements []*structureElement
	attributes        []string
}

const beginKeyWord = "BEGIN:"
const endKeyWord = "END:"

func readStructure(lines []string) ([]*structureElement, error) {
	read, err := readLines(lines, 0, "")
	if err != nil {
		return nil, err
	}

	if len(read.attributes) > 0 {
		return nil, fmt.Errorf("%d attributes lines found at the top level", len(read.attributes))
	}

	return read.structureElements, nil
}

func readLines(lines []string, fromIndex int, endLine string) (*readingAccumulator, error) {
	structureElements := make([]*structureElement, 0)
	attributes := make([]string, 0)

	currentIndex := fromIndex
	endLineFound := false

	for ; currentIndex < len(lines) && !endLineFound; currentIndex++ {
		line := lines[currentIndex]

		switch {
		case line == endLine:
			endLineFound = true

		case strings.HasPrefix(line, endKeyWord):
			return nil, fmt.Errorf("mismatched elements, expected \"%s\" and found \"%s\"", endLine, line)

		case strings.HasPrefix(line, beginKeyWord):
			structureElement, subElementEndIndex, err := readStructureElement(lines, currentIndex)
			if err != nil {
				return nil, err
			}

			structureElements = append(structureElements, structureElement)
			currentIndex = subElementEndIndex

		default:
			attributes = append(attributes, line)
		}
	}

	lastLineIndex := currentIndex - 1

	result := readingAccumulator{
		structureElements: structureElements,
		attributes:        attributes,
		lastLineIndex:     lastLineIndex,
	}

	return &result, nil
}

func readStructureElement(lines []string, fromIndex int) (*structureElement, int, error) {
	elementName := strings.TrimPrefix(lines[fromIndex], beginKeyWord)
	endLine := endKeyWord + elementName

	read, err := readLines(lines, fromIndex+1, endLine)
	if err != nil {
		return nil, -1, err
	}

	if lines[read.lastLineIndex] != endLine {
		return nil, -1, fmt.Errorf("unable to find the end of an structure element \"%s\"", elementName)
	}

	structureElement := structureElement{
		name:              elementName,
		structureElements: read.structureElements,
		attributes:        read.attributes,
	}

	return &structureElement, read.lastLineIndex, nil
}
