package folder

import (
	"bufio"
	"io"
	"strings"
)

// UnfoldLines reads a raw ics a return its content lines. During the reading, lines are unfolded. Also the ics folder
// breaks ("\n" or "\N") are converted to proper folder breaks
func UnfoldLines(reader io.Reader) ([]string, error) {
	lines := make([]string, 0)
	accumulator := ""

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		line = strings.ReplaceAll(line, "\\n", "\n")
		line = strings.ReplaceAll(line, "\\N", "\n")

		// Lines are folded if they start with space or tabulation
		if line[0] == ' ' || line[0] == '\t' {
			accumulator += line[1:]
		} else {
			if len(accumulator) > 0 {
				lines = append(lines, accumulator)
			}
			accumulator = line
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(accumulator) > 0 {
		lines = append(lines, accumulator)
	}

	return lines, nil
}
